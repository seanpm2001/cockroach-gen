// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package xform

import (
	"github.com/cockroachdb/cockroach/pkg/sql/opt"
	"github.com/cockroachdb/cockroach/pkg/sql/opt/cat"
	"github.com/cockroachdb/cockroach/pkg/sql/opt/memo"
	"github.com/cockroachdb/cockroach/pkg/sql/opt/ordering"
	"github.com/cockroachdb/cockroach/pkg/sql/opt/props"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
	"github.com/cockroachdb/cockroach/pkg/util"
	"github.com/cockroachdb/errors"
)

// IsCanonicalGroupBy returns true if the private is for the canonical version
// of the grouping operator. This is the operator that is built initially (and
// has all grouping columns as optional in the ordering), as opposed to variants
// generated by the GenerateStreamingGroupBy exploration rule.
func (c *CustomFuncs) IsCanonicalGroupBy(private *memo.GroupingPrivate) bool {
	return private.Ordering.Any() || private.GroupingCols.SubsetOf(private.Ordering.Optional)
}

// MakeMinMaxScalarSubqueries transforms a list of MIN and MAX aggregate
// expressions (aggs) and a scanPrivate into multiple scalar subqueries, with
// one MIN or MAX expression per subquery.
func (c *CustomFuncs) MakeMinMaxScalarSubqueries(
	grp memo.RelExpr, scanPrivate *memo.ScanPrivate, aggs memo.AggregationsExpr,
) {
	c.MakeMinMaxScalarSubqueriesWithFilter(grp, scanPrivate, aggs, nil)
}

// MakeMinMaxScalarSubqueriesWithFilter transforms a list of MIN and MAX aggregate
// expressions (aggs) and a scanPrivate with filters into multiple scalar
// subqueries, with one MIN or MAX expression per subquery.
func (c *CustomFuncs) MakeMinMaxScalarSubqueriesWithFilter(
	grp memo.RelExpr,
	scanPrivate *memo.ScanPrivate,
	aggs memo.AggregationsExpr,
	filters memo.FiltersExpr,
) {
	numCols := len(aggs)
	valuesEntries := make(memo.ScalarListExpr, numCols)
	dataTypes := make([]*types.T, numCols)
	newCols := make(opt.ColList, numCols)

	var inputExpr memo.RelExpr
	for i := 0; i < numCols; i++ {
		newScanPrivate := c.DuplicateScanPrivate(scanPrivate)
		inputExpr = c.e.f.ConstructScan(newScanPrivate)

		// If the input to the scalar group by is a Select with filters, remap the
		// column IDs in the filters and use that to build a new Select.
		if len(filters) > 0 {
			newFilters := c.RemapScanColsInFilter(filters, scanPrivate, newScanPrivate)
			inputExpr = c.e.f.ConstructSelect(inputExpr, newFilters)
		}

		var newAggrItem = aggs[i]
		// Variable expressions must have their ColIDs remapped, which requires
		// building a new AggregationsItem.
		variable, ok := aggs[i].Agg.Child(0).(*memo.VariableExpr)
		if !ok {
			panic(errors.AssertionFailedf("expected a variable as input to the aggregate, but found %T", aggs[i].Agg.Child(0)))
		}
		newVarExpr := c.remapScanColsInScalarExpr(variable, scanPrivate, newScanPrivate)
		var newAggrFunc opt.ScalarExpr
		switch aggs[i].Agg.(type) {
		case *memo.MaxExpr:
			newAggrFunc = c.e.f.ConstructMax(newVarExpr)
		case *memo.MinExpr:
			newAggrFunc = c.e.f.ConstructMin(newVarExpr)
		default:
			panic(errors.AssertionFailedf("expected a MIN or MAX expression, but found %T", aggs[i].Agg))
		}
		newAggrItem =
			c.e.f.ConstructAggregationsItem(
				newAggrFunc,
				aggs[i].Col,
			)

		valuesEntries[i] =
			c.e.f.ConstructSubquery(
				c.e.f.ConstructScalarGroupBy(
					inputExpr,
					memo.AggregationsExpr{
						newAggrItem,
					},
					memo.EmptyGroupingPrivate,
				),
				c.e.funcs.MakeUnorderedSubquery(),
			)
		// The VALUES clause inherits the target column ID and datatype from the
		// scalar group by, so the outputs are identical.
		dataTypes[i] = aggs[i].DataType()
		newCols[i] = aggs[i].Col
	}

	tupleTyp := types.MakeTuple(dataTypes)
	tuples := memo.ScalarListExpr{c.e.f.ConstructTuple(valuesEntries, tupleTyp)}

	var newPrivate = memo.ValuesPrivate{
		Cols: newCols,
		ID:   c.e.f.Metadata().NextUniqueID(),
	}
	valuesExpression := &memo.ValuesExpr{
		Rows:          tuples,
		ValuesPrivate: newPrivate,
	}
	c.e.mem.AddValuesToGroup(valuesExpression, grp)
}

// TwoOrMoreMinOrMax returns true if the aggregations (aggs) consists of two
// or more MIN or MAX expressions on variable expressions with no other type of
// aggregate function.
func (c *CustomFuncs) TwoOrMoreMinOrMax(aggs memo.AggregationsExpr) bool {
	if len(aggs) < 2 {
		return false
	}

	for i := range aggs {
		agg := &aggs[i]

		switch agg.Agg.Op() {
		case opt.MinOp, opt.MaxOp:
			// The child of a min or max aggregation should always be a variable, but
			// we add a sanity check here anyway.
			if _, ok := agg.Agg.Child(0).(*memo.VariableExpr); !ok {
				return false
			}
		default:
			return false
		}
	}
	return true
}

// MakeProjectFromPassthroughAggs constructs a top-level Project operator that
// contains one output column per function in the given aggregrate list. The
// input expression is expected to return zero or one rows, and the aggregate
// functions are expected to always pass through their values in that case.
func (c *CustomFuncs) MakeProjectFromPassthroughAggs(
	grp memo.RelExpr, input memo.RelExpr, aggs memo.AggregationsExpr,
) {
	if !input.Relational().Cardinality.IsZeroOrOne() {
		panic(errors.AssertionFailedf("input expression cannot have more than one row: %v", input))
	}

	var passthrough opt.ColSet
	projections := make(memo.ProjectionsExpr, 0, len(aggs))
	for i := range aggs {
		// If aggregate remaps the column ID, need to synthesize projection item;
		// otherwise, can just pass through.
		variable := aggs[i].Agg.Child(0).(*memo.VariableExpr)
		if variable.Col == aggs[i].Col {
			passthrough.Add(variable.Col)
		} else {
			projections = append(projections, c.e.f.ConstructProjectionsItem(variable, aggs[i].Col))
		}
	}
	c.e.mem.AddProjectToGroup(&memo.ProjectExpr{
		Input:       input,
		Projections: projections,
		Passthrough: passthrough,
	}, grp)
}

// GroupingColsClosureOverlappingOrdering determines if the specified `ordering`
// columns overlaps with the closure of the grouping columns in `private` as
// determined by the functional dependencies of the `input` relation to the
// grouped aggregation. If found, this expanded set of `groupingCols` is
// returned, plus the overlapping ordering columns `newOrdering`, along with
// `ok`=true.
func (c *CustomFuncs) GroupingColsClosureOverlappingOrdering(
	input memo.RelExpr, private *memo.GroupingPrivate, ordering props.OrderingChoice,
) (groupingCols opt.ColSet, newOrdering props.OrderingChoice, ok bool) {
	if ordering.Any() {
		// If the limit specifies no ordering, there is no ordering hint for us to
		// use.
		return opt.ColSet{}, props.OrderingChoice{}, false
	}
	groupingCols = private.GroupingCols
	// If the result requires a specific ordering, use that as the ordering of the
	// aggregation if possible.
	// Find all columns determined by the grouping columns.
	groupingColumnsClosure := input.Relational().FuncDeps.ComputeClosure(groupingCols)
	// It is safe to add ordering columns present in the grouping column closure
	// as grouping columns because the original grouping columns determine all
	// other column values in the closure (within a group there is only one
	// combined set of values for the other columns). Doing so allows the required
	// ordering to be provided by the streaming group by and possibly remove the
	// requirement for a TopK operator.
	orderingColsInClosure := groupingColumnsClosure.Intersection(ordering.ColSet())
	if orderingColsInClosure.Empty() {
		// If we have no columns to add to the grouping, this rewrite has no effect.
		return opt.ColSet{}, props.OrderingChoice{}, false
	}
	groupingCols = groupingCols.Union(orderingColsInClosure)
	newOrdering, fullPrefix, found := getPrefixFromOrdering(ordering.ToOrdering(), private.Ordering, input,
		func(id opt.ColumnID) bool { return groupingCols.Contains(id) })
	if !found || !fullPrefix {
		return opt.ColSet{}, props.OrderingChoice{}, false
	}
	return groupingCols, newOrdering, true
}

// GenerateStreamingGroupByLimitOrderingHint generates a LimitExpr with an input
// which is a GroupBy or DistinctOn expression with an expanded set of
// `groupingCols` which is an overlap of the closure of the original grouping
// columns closure and the ordering in the `limitExpr` as determined by
// `GroupingColsClosureOverlappingOrdering`, which also produces the
// `newOrdering`. Argument `private` is expected to be a canonical group-by.
func (c *CustomFuncs) GenerateStreamingGroupByLimitOrderingHint(
	grp memo.RelExpr,
	limitExpr *memo.LimitExpr,
	aggregation memo.RelExpr,
	input memo.RelExpr,
	aggs memo.AggregationsExpr,
	private *memo.GroupingPrivate,
	groupingCols opt.ColSet,
	newOrdering props.OrderingChoice,
) {
	if !c.e.evalCtx.SessionData().OptimizerUseLimitOrderingForStreamingGroupBy {
		// This transformation rule is explicitly disabled.
		return
	}
	newPrivate := *private
	newPrivate.Ordering = newOrdering
	newPrivate.GroupingCols = groupingCols

	// Remove constant column aggregate expressions if they've been added to
	// the grouping columns. The column should appear in one place or the other,
	// but not both.
	newAggs := make(memo.AggregationsExpr, 0, len(aggs))
	for _, agg := range aggs {
		if !groupingCols.Contains(agg.Col) {
			newAggs = append(newAggs, agg)
			continue
		}
		constAggExpr, ok := agg.Agg.(*memo.ConstAggExpr)
		if !ok {
			newAggs = append(newAggs, agg)
			continue
		}
		variableExpr, ok := constAggExpr.Input.(*memo.VariableExpr)
		if !ok {
			newAggs = append(newAggs, agg)
			continue
		}
		if variableExpr.Col != agg.Col {
			// Column IDs expected to match to safely remove this aggregation.
			return
		}
	}
	// Output columns are the union of grouping columns with columns from the
	// aggregate projection list. Verify this is built correctly.
	outputCols := groupingCols.Copy()
	for i := range newAggs {
		outputCols.Add(newAggs[i].Col)
	}
	if !aggregation.Relational().OutputCols.Equals(outputCols) {
		// If the output columns in the new aggregation don't match those in the
		// original aggregation, give up on this optimization.
		return
	}

	var newAggregation memo.RelExpr
	constructAggregation := func() {
		newAggregation =
			c.e.f.DynamicConstruct(
				aggregation.Op(),
				input,
				&newAggs,
				&newPrivate,
			).(memo.RelExpr)
	}
	var disabledRules util.FastIntSet
	// The ReduceGroupingCols rule must be disabled to prevent the ordering
	// columns from being removed from the grouping columns during operation
	// construction. This rule already reduced the grouping columns on the initial
	// construction. We are just adding back in any ordering columns which overlap
	// with grouping columns in order to generate a better plan.
	disabledRules.Add(int(opt.ReduceGroupingCols))

	c.e.f.DisableOptimizationRulesTemporarily(disabledRules, constructAggregation)
	newLimitExpr :=
		&memo.LimitExpr{
			Input:    newAggregation,
			Limit:    limitExpr.Limit,
			Ordering: limitExpr.Ordering,
		}
	grp.Memo().AddLimitToGroup(newLimitExpr, grp)
}

// GenerateStreamingGroupBy generates variants of a GroupBy, DistinctOn,
// EnsureDistinctOn, UpsertDistinctOn, or EnsureUpsertDistinctOn expression
// with more specific orderings on the grouping columns, using the interesting
// orderings property. See the GenerateStreamingGroupBy rule.
func (c *CustomFuncs) GenerateStreamingGroupBy(
	grp memo.RelExpr,
	op opt.Operator,
	input memo.RelExpr,
	aggs memo.AggregationsExpr,
	private *memo.GroupingPrivate,
) {
	orders := ordering.DeriveInterestingOrderings(input)
	intraOrd := private.Ordering
	for _, ord := range orders {
		newOrd, fullPrefix, found := getPrefixFromOrdering(ord.ToOrdering(), intraOrd, input,
			func(id opt.ColumnID) bool { return private.GroupingCols.Contains(id) })
		if !found || !fullPrefix {
			continue
		}

		newPrivate := *private
		newPrivate.Ordering = newOrd

		switch op {
		case opt.GroupByOp:
			newExpr := memo.GroupByExpr{
				Input:           input,
				Aggregations:    aggs,
				GroupingPrivate: newPrivate,
			}
			c.e.mem.AddGroupByToGroup(&newExpr, grp)

		case opt.DistinctOnOp:
			newExpr := memo.DistinctOnExpr{
				Input:           input,
				Aggregations:    aggs,
				GroupingPrivate: newPrivate,
			}
			c.e.mem.AddDistinctOnToGroup(&newExpr, grp)

		case opt.EnsureDistinctOnOp:
			newExpr := memo.EnsureDistinctOnExpr{
				Input:           input,
				Aggregations:    aggs,
				GroupingPrivate: newPrivate,
			}
			c.e.mem.AddEnsureDistinctOnToGroup(&newExpr, grp)

		case opt.UpsertDistinctOnOp:
			newExpr := memo.UpsertDistinctOnExpr{
				Input:           input,
				Aggregations:    aggs,
				GroupingPrivate: newPrivate,
			}
			c.e.mem.AddUpsertDistinctOnToGroup(&newExpr, grp)

		case opt.EnsureUpsertDistinctOnOp:
			newExpr := memo.EnsureUpsertDistinctOnExpr{
				Input:           input,
				Aggregations:    aggs,
				GroupingPrivate: newPrivate,
			}
			c.e.mem.AddEnsureUpsertDistinctOnToGroup(&newExpr, grp)
		}
	}
}

// OtherAggsAreConst returns true if all items in the given aggregate list
// contain ConstAgg functions, except for the "except" item. The ConstAgg
// functions will always return the same value, as long as there is at least
// one input row.
func (c *CustomFuncs) OtherAggsAreConst(
	aggs memo.AggregationsExpr, except *memo.AggregationsItem,
) bool {
	for i := range aggs {
		agg := &aggs[i]
		if agg == except {
			continue
		}

		switch agg.Agg.Op() {
		case opt.ConstAggOp:
			// Ensure that argument is a VariableOp.
			if agg.Agg.Child(0).Op() != opt.VariableOp {
				return false
			}

		default:
			return false
		}
	}
	return true
}

// MakeOrderingChoiceFromColumn constructs a new OrderingChoice with
// one element in the sequence: the columnID in the order defined by
// (MIN/MAX) operator. This function was originally created to be used
// with the Replace(Min|Max)WithLimit exploration rules.
//
// WARNING: The MinOp case can return a NULL value if the column allows it. This
// is because NULL values sort first in CRDB.
func (c *CustomFuncs) MakeOrderingChoiceFromColumn(
	op opt.Operator, col opt.ColumnID,
) props.OrderingChoice {
	oc := props.OrderingChoice{}
	switch op {
	case opt.MinOp:
		oc.AppendCol(col, false /* descending */)
	case opt.MaxOp:
		oc.AppendCol(col, true /* descending */)
	}
	return oc
}

// SplitGroupByScanIntoUnionScans splits a non-inverted scan under a GroupBy,
// DistinctOn, or EnsureUpsertDistinctOn into a UnionAll of scans, where each
// scan can provide an ordering on the grouping columns. If no such UnionAll
// can be built, returns ok=false.
//
// This is useful because the GenerateStreamingGroupBy rule can then create a
// streaming grouping operation, which is more efficient.
// GenerateStreamingGroupBy will use the new interesting orderings provided by
// the UnionAll of scans to build the streaming operation.
//
// See the SplitGroupByScanIntoUnionScans rule for more details.
func (c *CustomFuncs) SplitGroupByScanIntoUnionScans(
	scan memo.RelExpr, sp *memo.ScanPrivate, private *memo.GroupingPrivate,
) (_ memo.RelExpr, ok bool) {
	cons, ok := c.getKnownScanConstraint(sp)
	if !ok {
		// No valid constraint was found.
		return nil, false
	}

	intraOrd := private.Ordering

	// Find the length of the prefix of index columns preceding the first groupby
	// ordering column. We will verify later that the entire ordering sequence is
	// represented in the index. Ex:
	//
	//     Index: +1/+2/-3, Group By internal ordering +3 opt(4) => Prefix Length: 2
	//
	keyPrefixLength := cons.Columns.Count()
	for i := 0; i < cons.Columns.Count(); i++ {
		col := cons.Columns.Get(i).ID()
		if private.GroupingCols.Contains(col) || intraOrd.Optional.Contains(col) {
			// Grouping or optional column.
			keyPrefixLength = i
			break
		}
		if len(intraOrd.Columns) > 0 &&
			intraOrd.Group(0).Contains(col) {
			// Column matches the one in the ordering.
			keyPrefixLength = i
			break
		}
	}
	if keyPrefixLength == 0 {
		// This case can be handled by GenerateStreamingGroupBy.
		return nil, false
	}

	// Create a UnionAll of scans that can provide the ordering of the
	// GroupingPrivate (if no such UnionAll is possible this will return
	// ok=false). We pass a limit of 0 since the scans are unlimited
	// (splitScanIntoUnionScans is also used for another rule with limited scans).
	return c.splitScanIntoUnionScans(
		intraOrd, scan, sp, cons, 0 /* limit */, keyPrefixLength,
	)
}

// GroupingColumns returns the grouping columns from the grouping private.
func (c *CustomFuncs) GroupingColumns(private *memo.GroupingPrivate) opt.ColSet {
	return private.GroupingCols
}

// GroupingOrdering returns the ordering from the grouping private.
func (c *CustomFuncs) GroupingOrdering(private *memo.GroupingPrivate) props.OrderingChoice {
	return private.Ordering
}

// MakeGroupingPrivate constructs a new GroupingPrivate using the given
// grouping columns, OrderingChoice, NullsAreDistinct bool, and ErrorOnDup text.
func (c *CustomFuncs) MakeGroupingPrivate(
	groupingCols opt.ColSet, ordering props.OrderingChoice, nullsAreDistinct bool, errorText string,
) *memo.GroupingPrivate {
	return &memo.GroupingPrivate{
		GroupingCols:     groupingCols,
		Ordering:         ordering,
		NullsAreDistinct: nullsAreDistinct,
		ErrorOnDup:       errorText,
	}
}

// GenerateLimitedGroupByScans enumerates all non-inverted secondary indexes on
// the given Scan operator's table and generates an alternate Scan operator for
// each index that includes a partial set of needed columns specified in the
// ScanOpDef. An IndexJoin is constructed to add missing columns. A GroupBy and
// Limit are also constructed to make an equivalent expression for the memo.
//
// For cases where the Scan's secondary index covers all needed columns, see
// GenerateIndexScans, which does not construct an IndexJoin.
func (c *CustomFuncs) GenerateLimitedGroupByScans(
	grp memo.RelExpr,
	sp *memo.ScanPrivate,
	aggs memo.AggregationsExpr,
	gp *memo.GroupingPrivate,
	limit opt.ScalarExpr,
	required props.OrderingChoice,
) {
	// If the required ordering and grouping columns do not share columns, then
	// this optimization is not beneficial.
	if !required.Any() && !required.Group(0).Intersects(gp.GroupingCols) && !required.Optional.Intersects(gp.GroupingCols) {
		return
	}
	// Iterate over all non-inverted and non-partial secondary indexes.
	var pkCols opt.ColSet
	var iter scanIndexIter
	var sb indexScanBuilder
	sb.Init(c, sp.Table)
	iter.Init(c.e.evalCtx, c.e.f, c.e.mem, &c.im, sp, nil /* filters */, rejectPrimaryIndex|rejectInvertedIndexes)
	iter.ForEach(func(index cat.Index, filters memo.FiltersExpr, indexCols opt.ColSet, isCovering bool, constProj memo.ProjectionsExpr) {
		// The iterator only produces pseudo-partial indexes (the predicate is
		// true) because no filters are passed to iter.Init to imply a partial
		// index predicate. constProj is a projection of constant values based
		// on a partial index predicate. It should always be empty because a
		// pseudo-partial index cannot hold a column constant. If it is not, we
		// panic to avoid performing a logically incorrect transformation.
		if len(constProj) != 0 {
			panic(errors.AssertionFailedf("expected constProj to be empty"))
		}

		// If the secondary index includes the set of needed columns, then this
		// case does not need a limited group by and will be covered in
		// GenerateIndexScans.
		if isCovering {
			return
		}

		// Otherwise, try to construct an IndexJoin operator that provides the
		// columns missing from the index.
		if sp.Flags.NoIndexJoin {
			return
		}

		// Calculate the PK columns once.
		if pkCols.Empty() {
			pkCols = c.PrimaryKeyCols(sp.Table)
		}

		// If the first index column is not in the grouping columns, then there is
		// no benefit to exploring this index.
		if col := sp.Table.ColumnID(index.Column(0).Ordinal()); !gp.GroupingCols.Contains(col) {
			return
		}

		// If the index doesn't contain any of the required order columns, then
		// there is no benefit to exploring this index.
		if !required.Any() && !required.Group(0).Intersects(indexCols) {
			return
		}

		// Scan whatever columns we need which are available from the index.
		newScanPrivate := *sp
		newScanPrivate.Index = index.Ordinal()
		newScanPrivate.Cols = indexCols.Intersection(sp.Cols)
		// If the index is not covering, scan the needed index columns plus
		// primary key columns.
		newScanPrivate.Cols.UnionWith(pkCols)
		sb.SetScan(&newScanPrivate)
		// Construct an IndexJoin operator that provides the columns missing from
		// the index.
		sb.AddIndexJoin(sp.Cols)
		input := sb.BuildNewExpr()
		// Reconstruct the GroupBy and Limit so the new expression in the memo is
		// equivalent.
		input = c.e.f.ConstructGroupBy(input, aggs, gp)
		grp.Memo().AddLimitToGroup(&memo.LimitExpr{Limit: limit, Ordering: required, Input: input}, grp)
	})
}
