// Code generated by optgen; DO NOT EDIT.

package opt

type Factory interface {
	// Metadata returns the query-specific metadata, which includes information
	// about the columns and tables used in this particular query.
	Metadata() *Metadata

	// InternList adds the given list of group IDs to memo storage and returns
	// an ID that can be used for later lookup. If the same list was added
	// previously, this method is a no-op and returns the ID of the previous
	// value.
	InternList(items []GroupID) ListID

	// InternPrivate adds the given private value to the memo and returns an ID
	// that can be used for later lookup. If the same value was added before,
	// then this method is a no-op and returns the ID of the previous value.
	InternPrivate(private interface{}) PrivateID

	// DynamicConstruct dynamically constructs an operator with the given type
	// and operands. It is equivalent to a switch statement that calls the
	// ConstructXXX method that corresponds to the given operator.
	DynamicConstruct(op Operator, children []GroupID, private PrivateID) GroupID

	// ------------------------------------------------------------
	// Scalar Operators
	// ------------------------------------------------------------

	// ConstructSubquery constructs an expression for the Subquery operator.
	ConstructSubquery(input GroupID, projection GroupID) GroupID

	// ConstructVariable constructs an expression for the Variable operator.
	// Variable is the typed scalar value of a column in the query. The private
	// field is a Metadata.ColumnIndex that references the column by index.
	ConstructVariable(col PrivateID) GroupID

	// ConstructConst constructs an expression for the Const operator.
	// Const is a typed scalar constant value. The private field is a tree.Datum
	// value having any datum type that's legal in the expression's context.
	ConstructConst(value PrivateID) GroupID

	// ConstructTrue constructs an expression for the True operator.
	// True is the boolean true value that is equivalent to the tree.DBoolTrue datum
	// value. It is a separate operator to make matching and replacement simpler and
	// more efficient, as patterns can contain (True) expressions.
	ConstructTrue() GroupID

	// ConstructFalse constructs an expression for the False operator.
	// False is the boolean false value that is equivalent to the tree.DBoolFalse
	// datum value. It is a separate operator to make matching and replacement
	// simpler and more efficient, as patterns can contain (False) expressions.
	ConstructFalse() GroupID

	// ConstructPlaceholder constructs an expression for the Placeholder operator.
	ConstructPlaceholder(value PrivateID) GroupID

	// ConstructTuple constructs an expression for the Tuple operator.
	ConstructTuple(elems ListID) GroupID

	// ConstructProjections constructs an expression for the Projections operator.
	// Projections is a set of typed scalar expressions that will become output
	// columns for a containing Project operator. The private Cols field contains
	// the list of column indexes returned by the expression, as a *opt.ColList. It
	// is not legal for Cols to be empty.
	ConstructProjections(elems ListID, cols PrivateID) GroupID

	// ConstructAggregations constructs an expression for the Aggregations operator.
	// Aggregations is a set of aggregate expressions that will become output
	// columns for a containing GroupBy operator. The private Cols field contains
	// the list of column indexes returned by the expression, as a *ColList. It
	// is legal for Cols to be empty.
	ConstructAggregations(aggs ListID, cols PrivateID) GroupID

	// ConstructExists constructs an expression for the Exists operator.
	ConstructExists(input GroupID) GroupID

	// ConstructAnd constructs an expression for the And operator.
	// And is the boolean conjunction operator that evalutes to true if all of its
	// conditions evaluate to true. If the conditions list is empty, it evalutes to
	// true.
	ConstructAnd(conditions ListID) GroupID

	// ConstructOr constructs an expression for the Or operator.
	// Or is the boolean disjunction operator that evalutes to true if any of its
	// conditions evaluate to true. If the conditions list is empty, it evaluates to
	// false.
	ConstructOr(conditions ListID) GroupID

	// ConstructNot constructs an expression for the Not operator.
	// Not is the boolean negation operator that evaluates to true if its input
	// evalutes to false.
	ConstructNot(input GroupID) GroupID

	// ConstructEq constructs an expression for the Eq operator.
	ConstructEq(left GroupID, right GroupID) GroupID

	// ConstructLt constructs an expression for the Lt operator.
	ConstructLt(left GroupID, right GroupID) GroupID

	// ConstructGt constructs an expression for the Gt operator.
	ConstructGt(left GroupID, right GroupID) GroupID

	// ConstructLe constructs an expression for the Le operator.
	ConstructLe(left GroupID, right GroupID) GroupID

	// ConstructGe constructs an expression for the Ge operator.
	ConstructGe(left GroupID, right GroupID) GroupID

	// ConstructNe constructs an expression for the Ne operator.
	ConstructNe(left GroupID, right GroupID) GroupID

	// ConstructIn constructs an expression for the In operator.
	ConstructIn(left GroupID, right GroupID) GroupID

	// ConstructNotIn constructs an expression for the NotIn operator.
	ConstructNotIn(left GroupID, right GroupID) GroupID

	// ConstructLike constructs an expression for the Like operator.
	ConstructLike(left GroupID, right GroupID) GroupID

	// ConstructNotLike constructs an expression for the NotLike operator.
	ConstructNotLike(left GroupID, right GroupID) GroupID

	// ConstructILike constructs an expression for the ILike operator.
	ConstructILike(left GroupID, right GroupID) GroupID

	// ConstructNotILike constructs an expression for the NotILike operator.
	ConstructNotILike(left GroupID, right GroupID) GroupID

	// ConstructSimilarTo constructs an expression for the SimilarTo operator.
	ConstructSimilarTo(left GroupID, right GroupID) GroupID

	// ConstructNotSimilarTo constructs an expression for the NotSimilarTo operator.
	ConstructNotSimilarTo(left GroupID, right GroupID) GroupID

	// ConstructRegMatch constructs an expression for the RegMatch operator.
	ConstructRegMatch(left GroupID, right GroupID) GroupID

	// ConstructNotRegMatch constructs an expression for the NotRegMatch operator.
	ConstructNotRegMatch(left GroupID, right GroupID) GroupID

	// ConstructRegIMatch constructs an expression for the RegIMatch operator.
	ConstructRegIMatch(left GroupID, right GroupID) GroupID

	// ConstructNotRegIMatch constructs an expression for the NotRegIMatch operator.
	ConstructNotRegIMatch(left GroupID, right GroupID) GroupID

	// ConstructIs constructs an expression for the Is operator.
	ConstructIs(left GroupID, right GroupID) GroupID

	// ConstructIsNot constructs an expression for the IsNot operator.
	ConstructIsNot(left GroupID, right GroupID) GroupID

	// ConstructContains constructs an expression for the Contains operator.
	ConstructContains(left GroupID, right GroupID) GroupID

	// ConstructBitand constructs an expression for the Bitand operator.
	ConstructBitand(left GroupID, right GroupID) GroupID

	// ConstructBitor constructs an expression for the Bitor operator.
	ConstructBitor(left GroupID, right GroupID) GroupID

	// ConstructBitxor constructs an expression for the Bitxor operator.
	ConstructBitxor(left GroupID, right GroupID) GroupID

	// ConstructPlus constructs an expression for the Plus operator.
	ConstructPlus(left GroupID, right GroupID) GroupID

	// ConstructMinus constructs an expression for the Minus operator.
	ConstructMinus(left GroupID, right GroupID) GroupID

	// ConstructMult constructs an expression for the Mult operator.
	ConstructMult(left GroupID, right GroupID) GroupID

	// ConstructDiv constructs an expression for the Div operator.
	ConstructDiv(left GroupID, right GroupID) GroupID

	// ConstructFloorDiv constructs an expression for the FloorDiv operator.
	ConstructFloorDiv(left GroupID, right GroupID) GroupID

	// ConstructMod constructs an expression for the Mod operator.
	ConstructMod(left GroupID, right GroupID) GroupID

	// ConstructPow constructs an expression for the Pow operator.
	ConstructPow(left GroupID, right GroupID) GroupID

	// ConstructConcat constructs an expression for the Concat operator.
	ConstructConcat(left GroupID, right GroupID) GroupID

	// ConstructLShift constructs an expression for the LShift operator.
	ConstructLShift(left GroupID, right GroupID) GroupID

	// ConstructRShift constructs an expression for the RShift operator.
	ConstructRShift(left GroupID, right GroupID) GroupID

	// ConstructFetchVal constructs an expression for the FetchVal operator.
	ConstructFetchVal(json GroupID, index GroupID) GroupID

	// ConstructFetchText constructs an expression for the FetchText operator.
	ConstructFetchText(json GroupID, index GroupID) GroupID

	// ConstructFetchValPath constructs an expression for the FetchValPath operator.
	ConstructFetchValPath(json GroupID, path GroupID) GroupID

	// ConstructFetchTextPath constructs an expression for the FetchTextPath operator.
	ConstructFetchTextPath(json GroupID, path GroupID) GroupID

	// ConstructUnaryPlus constructs an expression for the UnaryPlus operator.
	ConstructUnaryPlus(input GroupID) GroupID

	// ConstructUnaryMinus constructs an expression for the UnaryMinus operator.
	ConstructUnaryMinus(input GroupID) GroupID

	// ConstructUnaryComplement constructs an expression for the UnaryComplement operator.
	ConstructUnaryComplement(input GroupID) GroupID

	// ConstructCast constructs an expression for the Cast operator.
	ConstructCast(input GroupID, typ PrivateID) GroupID

	// ConstructFunction constructs an expression for the Function operator.
	// Function invokes a builtin SQL function like CONCAT or NOW, passing the given
	// arguments. The private field is an opt.FuncDef struct that provides the name
	// of the function as well as a pointer to the builtin overload definition.
	ConstructFunction(args ListID, def PrivateID) GroupID

	// ConstructCoalesce constructs an expression for the Coalesce operator.
	ConstructCoalesce(args ListID) GroupID

	// ConstructUnsupportedExpr constructs an expression for the UnsupportedExpr operator.
	// UnsupportedExpr is used for interfacing with the old planner code. It can
	// encapsulate a TypedExpr that is otherwise not supported by the optimizer.
	ConstructUnsupportedExpr(value PrivateID) GroupID

	// ------------------------------------------------------------
	// Relational Operators
	// ------------------------------------------------------------

	// ConstructScan constructs an expression for the Scan operator.
	// Scan returns a result set containing every row in the specified table. Rows
	// and columns are not expected to have any particular ordering. The private
	// Table field is a Metadata.TableIndex that references an optbase.Table
	// definition in the query's metadata.
	ConstructScan(table PrivateID) GroupID

	// ConstructValues constructs an expression for the Values operator.
	// Values returns a manufactured result set containing a constant number of rows.
	// specified by the Rows list field. Each row must contain the same set of
	// columns in the same order.
	//
	// The Rows field contains a list of Tuples, one for each row. Each tuple has
	// the same length (same with that of Cols).
	//
	// The Cols field contains the set of column indices returned by each row
	// as a *ColList. It is legal for Cols to be empty.
	ConstructValues(rows ListID, cols PrivateID) GroupID

	// ConstructSelect constructs an expression for the Select operator.
	// Select filters rows from its input result set, based on the boolean filter
	// predicate expression. Rows which do not match the filter are discarded.
	ConstructSelect(input GroupID, filter GroupID) GroupID

	// ConstructProject constructs an expression for the Project operator.
	// Project modifies the set of columns returned by the input result set. Columns
	// can be removed, reordered, or renamed. In addition, new columns can be
	// synthesized. Projections is a scalar Projections list operator that contains
	// the list of expressions that describe the output columns. The Cols field of
	// the Projections operator provides the indexes of each of the output columns.
	ConstructProject(input GroupID, projections GroupID) GroupID

	// ConstructInnerJoin constructs an expression for the InnerJoin operator.
	// InnerJoin creates a result set that combines columns from its left and right
	// inputs, based upon its "on" join predicate. Rows which do not match the
	// predicate are filtered. While expressions in the predicate can refer to
	// columns projected by either the left or right inputs, the inputs are not
	// allowed to refer to the other's projected columns.
	ConstructInnerJoin(left GroupID, right GroupID, on GroupID) GroupID

	// ConstructLeftJoin constructs an expression for the LeftJoin operator.
	ConstructLeftJoin(left GroupID, right GroupID, on GroupID) GroupID

	// ConstructRightJoin constructs an expression for the RightJoin operator.
	ConstructRightJoin(left GroupID, right GroupID, on GroupID) GroupID

	// ConstructFullJoin constructs an expression for the FullJoin operator.
	ConstructFullJoin(left GroupID, right GroupID, on GroupID) GroupID

	// ConstructSemiJoin constructs an expression for the SemiJoin operator.
	ConstructSemiJoin(left GroupID, right GroupID, on GroupID) GroupID

	// ConstructAntiJoin constructs an expression for the AntiJoin operator.
	ConstructAntiJoin(left GroupID, right GroupID, on GroupID) GroupID

	// ConstructInnerJoinApply constructs an expression for the InnerJoinApply operator.
	// InnerJoinApply has the same join semantics as InnerJoin. However, unlike
	// InnerJoin, it allows the right input to refer to columns projected by the
	// left input.
	ConstructInnerJoinApply(left GroupID, right GroupID, on GroupID) GroupID

	// ConstructLeftJoinApply constructs an expression for the LeftJoinApply operator.
	ConstructLeftJoinApply(left GroupID, right GroupID, on GroupID) GroupID

	// ConstructRightJoinApply constructs an expression for the RightJoinApply operator.
	ConstructRightJoinApply(left GroupID, right GroupID, on GroupID) GroupID

	// ConstructFullJoinApply constructs an expression for the FullJoinApply operator.
	ConstructFullJoinApply(left GroupID, right GroupID, on GroupID) GroupID

	// ConstructSemiJoinApply constructs an expression for the SemiJoinApply operator.
	ConstructSemiJoinApply(left GroupID, right GroupID, on GroupID) GroupID

	// ConstructAntiJoinApply constructs an expression for the AntiJoinApply operator.
	ConstructAntiJoinApply(left GroupID, right GroupID, on GroupID) GroupID

	// ConstructGroupBy constructs an expression for the GroupBy operator.
	// GroupBy is an operator that is used for performing aggregations (for queries
	// with aggregate functions, HAVING clauses and/or group by expressions). It
	// groups results that are equal on the grouping columns and computes
	// aggregations as described by Aggregations (which is always an Aggregations
	// operator). The arguments of the aggregations are columns from the input.
	ConstructGroupBy(input GroupID, aggregations GroupID, groupingColumns PrivateID) GroupID

	// ConstructUnion constructs an expression for the Union operator.
	ConstructUnion(left GroupID, right GroupID, colMap PrivateID) GroupID

	// ConstructIntersect constructs an expression for the Intersect operator.
	ConstructIntersect(left GroupID, right GroupID) GroupID

	// ConstructExcept constructs an expression for the Except operator.
	ConstructExcept(left GroupID, right GroupID) GroupID
}
