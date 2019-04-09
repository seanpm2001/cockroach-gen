// Code generated by optgen; DO NOT EDIT.

package opt

const (
	startAutoRule RuleName = iota + NumManualRuleNames

	// ------------------------------------------------------------
	// Normalize Rule Names
	// ------------------------------------------------------------
	EliminateAggDistinct
	NormalizeNestedAnds
	SimplifyTrueAnd
	SimplifyAndTrue
	SimplifyFalseAnd
	SimplifyAndFalse
	SimplifyTrueOr
	SimplifyOrTrue
	SimplifyFalseOr
	SimplifyOrFalse
	SimplifyRange
	FoldNullAndOr
	FoldNotTrue
	FoldNotFalse
	NegateComparison
	EliminateNot
	NegateAnd
	NegateOr
	ExtractRedundantConjunct
	CommuteVarInequality
	CommuteConstInequality
	NormalizeCmpPlusConst
	NormalizeCmpMinusConst
	NormalizeCmpConstMinus
	NormalizeTupleEquality
	FoldNullComparisonLeft
	FoldNullComparisonRight
	FoldIsNull
	FoldNonNullIsNull
	FoldIsNotNull
	FoldNonNullIsNotNull
	CommuteNullIs
	DecorrelateJoin
	DecorrelateProjectSet
	TryDecorrelateSelect
	TryDecorrelateProject
	TryDecorrelateProjectSelect
	TryDecorrelateProjectInnerJoin
	TryDecorrelateInnerJoin
	TryDecorrelateInnerLeftJoin
	TryDecorrelateGroupBy
	TryDecorrelateScalarGroupBy
	TryDecorrelateSemiJoin
	TryDecorrelateLimitOne
	TryDecorrelateProjectSet
	HoistSelectExists
	HoistSelectNotExists
	HoistSelectSubquery
	HoistProjectSubquery
	HoistJoinSubquery
	HoistValuesSubquery
	HoistProjectSetSubquery
	NormalizeSelectAnyFilter
	NormalizeJoinAnyFilter
	NormalizeSelectNotAnyFilter
	NormalizeJoinNotAnyFilter
	FoldArray
	FoldBinary
	FoldUnary
	FoldComparison
	FoldCast
	FoldFunction
	ConvertGroupByToDistinct
	EliminateDistinct
	EliminateGroupByProject
	ReduceGroupingCols
	EliminateAggDistinctForKeys
	EliminateDistinctOnNoColumns
	InlineProjectConstants
	InlineSelectConstants
	InlineJoinConstantsLeft
	InlineJoinConstantsRight
	PushSelectIntoInlinableProject
	InlineProjectInProject
	SimplifyJoinFilters
	DetectJoinContradiction
	PushFilterIntoJoinLeftAndRight
	MapFilterIntoJoinLeft
	MapFilterIntoJoinRight
	PushFilterIntoJoinLeft
	PushFilterIntoJoinRight
	SimplifyLeftJoinWithoutFilters
	SimplifyRightJoinWithoutFilters
	SimplifyLeftJoinWithFilters
	SimplifyRightJoinWithFilters
	EliminateSemiJoin
	EliminateAntiJoin
	EliminateJoinNoColsLeft
	EliminateJoinNoColsRight
	HoistJoinProjectRight
	HoistJoinProjectLeft
	SimplifyJoinNotNullEquality
	ExtractJoinEqualities
	SortFiltersInJoin
	EliminateLimit
	PushLimitIntoProject
	PushOffsetIntoProject
	EliminateMax1Row
	FoldPlusZero
	FoldZeroPlus
	FoldMinusZero
	FoldMultOne
	FoldOneMult
	FoldDivOne
	InvertMinus
	EliminateUnaryMinus
	SimplifyLimitOrdering
	SimplifyOffsetOrdering
	SimplifyGroupByOrdering
	SimplifyRowNumberOrdering
	SimplifyExplainOrdering
	EliminateProject
	MergeProjects
	MergeProjectWithValues
	PruneProjectCols
	PruneScanCols
	PruneSelectCols
	PruneLimitCols
	PruneOffsetCols
	PruneJoinLeftCols
	PruneJoinRightCols
	PruneAggCols
	PruneGroupByCols
	PruneValuesCols
	PruneRowNumberCols
	PruneExplainCols
	PruneProjectSetCols
	PruneMutationFetchCols
	PruneMutationInputCols
	RejectNullsLeftJoin
	RejectNullsRightJoin
	RejectNullsGroupBy
	CommuteVar
	CommuteConst
	EliminateCoalesce
	SimplifyCoalesce
	EliminateCast
	FoldNullCast
	FoldNullUnary
	FoldNullBinaryLeft
	FoldNullBinaryRight
	FoldNullInNonEmpty
	FoldNullInEmpty
	FoldNullNotInEmpty
	NormalizeInConst
	FoldInNull
	UnifyComparisonTypes
	EliminateExistsProject
	EliminateExistsGroupBy
	IntroduceExistsLimit
	NormalizeJSONFieldAccess
	NormalizeJSONContains
	SimplifyCaseWhenConstValue
	SimplifyEqualsAnyTuple
	SimplifyAnyScalarArray
	FoldCollate
	NormalizeArrayFlattenToAgg
	SimplifySelectFilters
	ConsolidateSelectFilters
	DetectSelectContradiction
	EliminateSelect
	MergeSelects
	PushSelectIntoProject
	MergeSelectInnerJoin
	PushSelectCondLeftIntoJoinLeftAndRight
	PushSelectCondRightIntoJoinLeftAndRight
	PushSelectIntoJoinLeft
	PushSelectIntoJoinRight
	PushSelectIntoGroupBy
	RemoveNotNullCondition
	EliminateUnionAllLeft
	EliminateUnionAllRight

	// startExploreRule tracks the number of normalization rules;
	// all rules greater than this value are exploration rules.
	startExploreRule

	// ------------------------------------------------------------
	// Explore Rule Names
	// ------------------------------------------------------------
	ReplaceMinWithLimit
	ReplaceMaxWithLimit
	GenerateStreamingGroupBy
	CommuteJoin
	CommuteLeftJoin
	CommuteRightJoin
	GenerateMergeJoins
	GenerateLookupJoins
	GenerateZigzagJoins
	GenerateInvertedIndexZigzagJoins
	GenerateLookupJoinsWithFilter
	AssociateJoin
	GenerateLimitedScans
	PushLimitIntoConstrainedScan
	PushLimitIntoIndexJoin
	GenerateIndexScans
	GenerateConstrainedScans
	GenerateInvertedIndexScans

	// NumRuleNames tracks the total count of rule names.
	NumRuleNames
)
