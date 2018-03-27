// Copyright 2018 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package xform

import (
	"github.com/cockroachdb/cockroach/pkg/sql/opt"
	"github.com/cockroachdb/cockroach/pkg/sql/opt/memo"
	"github.com/cockroachdb/cockroach/pkg/util"
)

// explorer generates alternate expressions that are logically equivalent to
// existing expressions in the memo. The new expressions are added to the same
// memo group as the existing expression. The optimizer will cost all the
// expressions and pick the lowest cost alternative that provides any required
// physical properties.
//
// Equivalent expressions are generated by exploration rules. An exploration
// rule efficiently enumerates all possible combinations of its sub-expressions
// in order to look for matches. For example:
//
//  // [AssociateJoin]
//  (InnerJoin
//    (InnerJoin $r:* $s:* $lowerOn:*)
//    $t:*
//    $upperOn:*
//  )
//  =>
//  ...
//
// Say the memo group containing the upper inner-join has 3 expressions in it,
// and the memo group containing the lower inner-join has 4 expressions. Then
// the explorer will enumerate 12 possible expression combinations, looking for
// a combination that has an InnerJoin operator with another InnerJoin operator
// as its left operand.
//
// Once new expressions have been added to a group, they themselves become
// eligible for exploration, which might generate further expressions, and so
// on. Because the same group will often be explored multiple times, the
// explorer keeps state which helps it avoid duplicate work during subsequent
// passes.
//
// The explorer only traverses expression trees to the depth required by the
// exploration match patterns. It expects the optimizer to call exploreGroup
// for each group that needs to be explored. The optimizer can then use branch
// and bound pruning to skip exploration of entire sub-trees.
//
// For each expression combination that matches, a replace expression is
// constructed and added to the same memo group as the matched expression:
//
//  // [AssociateJoin]
//  (InnerJoin
//    (InnerJoin $r:* $s:* $lowerOn:*)
//    $t:*
//    $upperOn:*
//  )
//  =>
//  (InnerJoin
//    (InnerJoin
//      $r
//      $t
//      (Filters (ConstructConditionsNotUsing $s $lowerOn $upperOn))
//    )
//    $s
//    (Filters (ConstructConditionsUsing $s $lowerOn $upperOn))
//  )
//
// In this example, if the upper and lower groups each contain two InnerJoin
// expressions, then four new expressions will be added to the memo group of the
// matched expression. During the next pass, the four new expressions will
// themselves match this same rule. However, adding their replace expressions to
// the memo group will be a no-op, because they're already present.
type explorer struct {
	o   *Optimizer
	mem *memo.Memo
	f   *Factory
}

func (e *explorer) init(o *Optimizer) {
	e.o = o
	e.mem = o.mem
	e.f = o.f
}

// exploreGroup generates alternate expressions that are logically equivalent
// to the expressions already in the given group, and adds them to the group.
// The explorer maintains state that tracks which expressions were explored in
// previous passes. It keeps "start" and "end" expressions for the group which
// track the expressions which need to be fully explored during the current
// pass. Each time exploreGroup is called, the end of the previous pass becomes
// the start of the next pass. For example:
//
//   pass1         pass2         pass3
//      <-start
//   e0            e0            e0
//      <-end         <-start
//   e1 (new)      e1            e1
//
//   e2 (new)      e2            e2
//                    <-end         <-start
//                 e3 (new)      e3
//                                  <-end
//
// For rules which match one or more sub-expressions in addition to the top-
// level expression, there is extra complexity because every combination needs
// to be considered. Even expressions which were explored in previous passes
// need to be partially re-explored, because they may match when considered in
// combination with a new sub-expression which wasn't present during the last
// pass. Only combinations which consist solely of old expressions can be
// skipped.
//
// Combination enumeration code is just a series of nested loops generated by
// Optgen. Each non-scalar match pattern or sub-pattern uses a loop to
// enumerate the expressions in the corresponding memo group. For example:
//
//   $join:(InnerJoin
//     $left:(InnerJoin)
//     $right:(Select)
//     $on:*
//   )
//
// This match pattern would be implemented with 3 nested loops: 1 each for the
// $join, $left, and $right memo groups. If $join had 2 expressions, $left had
// 3 expressions, and $right had 2 expressions, then 2 * 3 * 2 = 12 combos will
// be considered. The innermost loop can skip iteration if all outer loops are
// bound to expressions which have already been explored in previous passes:
//
//   for e1 in memo-exprs($join):
//     for e2 in memo-exprs($left):
//       for e3 in memo-exprs($right):
//         if ordinal(e3) >= state.start:
//           ... explore (e1, e2, e3) combo ...
//
func (e *explorer) exploreGroup(group memo.GroupID) *exploreState {
	// Do nothing if this group has already been fully explored.
	state := e.ensureExploreState(group)
	if state.fullyExplored {
		return state
	}

	// Update set of expressions that will be considered during this pass, by
	// setting the start expression to be the end expression from last pass.
	exprCount := e.mem.ExprCount(group)
	state.start = state.end
	state.end = memo.ExprOrdinal(exprCount)

	fullyExplored := true
	for i := 0; i < exprCount; i++ {
		ordinal := memo.ExprOrdinal(i)

		// If expression was fully explored in previous passes, then nothing
		// further to do.
		if state.isExprFullyExplored(ordinal) {
			continue
		}

		eid := memo.ExprID{Group: group, Expr: ordinal}
		if e.exploreExpr(state, eid) {
			// No more rules can ever match this expression, so skip it in
			// future passes.
			state.markExprAsFullyExplored(ordinal)
		} else {
			// If even one expression is not fully explored, then the group is
			// not fully explored.
			fullyExplored = false
		}
	}

	// If all existing group expressions have been fully explored, and no new
	// ones were added, then group can be skipped in future passes.
	if fullyExplored && e.mem.ExprCount(group) == int(state.end) {
		state.fullyExplored = true
	}
	return state
}

// TODO(andyk): This dispatcher method needs to be generated by Optgen.
func (e *explorer) exploreExpr(state *exploreState, eid memo.ExprID) (fullyExplored bool) {
	switch e.mem.Expr(eid).Operator() {
	case opt.ScanOp:
		return e.exploreScan(state, eid)
	}

	// No rules for other operator types.
	return true
}

// TODO(andyk): This method needs to be generated by Optgen.
func (e *explorer) exploreScan(scanState *exploreState, scan memo.ExprID) (fullyExplored bool) {
	fullyExplored = true

	//	[GenerateIndexScans, Explore]
	//	$scan:(Scan & (IsPrimaryScan $scan)) => (ImplementIndexScan $scan)
	{
		// No sub-patterns, so this rule only needs to be matched once.
		if scan.Expr >= scanState.start {
			if e.isPrimaryScan(scan) {
				if e.o.allowOptimizations() {
					e.o.reportOptimization(GenerateIndexScans)
					e.generateIndexScans(scan)
				}
			}
		}
	}

	return fullyExplored
}

// ensureExploreState allocates the exploration state in the optState struct
// associated with the memo group, with respect to the min physical props.
func (e *explorer) ensureExploreState(group memo.GroupID) *exploreState {
	return &e.o.ensureOptState(group, memo.MinPhysPropsID).explore
}

// ----------------------------------------------------------------------
//
// Scan Rules
//   Custom match and replace functions used with scan.opt rules.
//
// ----------------------------------------------------------------------

// isPrimaryScan returns true if the given expression is scanning a primary
// index rather than a secondary index.
func (e *explorer) isPrimaryScan(scan memo.ExprID) bool {
	def := e.mem.LookupPrivate(e.mem.Expr(scan).AsScan().Def()).(*memo.ScanOpDef)
	return def.Index == opt.PrimaryIndex
}

// generateIndexScans enumerates all indexes on the scan operator's table and
// generates an alternate scan operator for each index that includes the set of
// needed columns.
// TODO(andyk): Create join with primary index in non-covering case.
func (e *explorer) generateIndexScans(scan memo.ExprID) {
	def := e.mem.LookupPrivate(e.mem.Expr(scan).AsScan().Def()).(*memo.ScanOpDef)
	tab := e.mem.Metadata().Table(def.Table)

	// Iterate over all secondary indexes (index 0 is the primary index).
	for i := 1; i < tab.IndexCount(); i++ {
		// If the alternate index includes the set of needed columns (def.Cols),
		// then construct a new Scan operator using that index.
		if def.AltIndexHasCols(e.mem.Metadata(), i) {
			newDef := &memo.ScanOpDef{Table: def.Table, Index: i, Cols: def.Cols}
			indexScan := memo.MakeScanExpr(e.mem.InternScanOpDef(newDef))
			e.mem.MemoizeDenormExpr(scan.Group, memo.Expr(indexScan))
		}
	}
}

// ----------------------------------------------------------------------
//
// Exploration state
//
// ----------------------------------------------------------------------

// exploreState contains state needed by the explorer for each memo group it
// explores. The state is allocated lazily for a group when the exploreGroup
// method is called. Various fields record what exploration has taken place so
// that the same work isn't repeated.
type exploreState struct {
	// start (inclusive) and end (exclusive) specify which expressions need to
	// be explored in the current pass. Expressions < start have been partly
	// explored during previous passes. Expressions >= end are new expressions
	// added during the current pass.
	start memo.ExprOrdinal
	end   memo.ExprOrdinal

	// fullyExplored is set to true once all expressions in the group have been
	// fully explored, and no new expressions will ever be added. Further
	// exploration of the group can be skipped.
	fullyExplored bool

	// fullyExploredExprs is a set of memo.ExprOrdinal values. Once an
	// expression has been fully explored, its ordinal is added to this set.
	fullyExploredExprs util.FastIntSet
}

// isExprFullyExplored is true if the given expression will never match an
// additional rule, and can therefore be skipped in future exploration passes.
func (e *exploreState) isExprFullyExplored(ordinal memo.ExprOrdinal) bool {
	return e.fullyExploredExprs.Contains(int(ordinal))
}

// markExprAsFullyExplored is called when all possible matching combinations
// have been considered for the subtree rooted at the given expression. Even if
// there are more exploration passes, this expression will never have new
// children, grand-children, etc. that might cause it to match another rule.
func (e *exploreState) markExprAsFullyExplored(ordinal memo.ExprOrdinal) {
	e.fullyExploredExprs.Add(int(ordinal))
}
