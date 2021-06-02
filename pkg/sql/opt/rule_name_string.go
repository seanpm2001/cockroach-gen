// Code generated by "stringer -output=pkg/sql/opt/rule_name_string.go -type=RuleName pkg/sql/opt/rule_name.go pkg/sql/opt/rule_name.og.go"; DO NOT EDIT.

package opt

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[InvalidRuleName-0]
	_ = x[SimplifyRootOrdering-1]
	_ = x[PruneRootCols-2]
	_ = x[SimplifyZeroCardinalityGroup-3]
	_ = x[NumManualRuleNames-4]
	_ = x[startAutoRule-4]
	_ = x[EliminateAggDistinct-5]
	_ = x[NormalizeNestedAnds-6]
	_ = x[SimplifyTrueAnd-7]
	_ = x[SimplifyAndTrue-8]
	_ = x[SimplifyFalseAnd-9]
	_ = x[SimplifyAndFalse-10]
	_ = x[SimplifyTrueOr-11]
	_ = x[SimplifyOrTrue-12]
	_ = x[SimplifyFalseOr-13]
	_ = x[SimplifyOrFalse-14]
	_ = x[SimplifyRange-15]
	_ = x[FoldNullAndOr-16]
	_ = x[FoldNotTrue-17]
	_ = x[FoldNotFalse-18]
	_ = x[FoldNotNull-19]
	_ = x[NegateComparison-20]
	_ = x[EliminateNot-21]
	_ = x[NegateAnd-22]
	_ = x[NegateOr-23]
	_ = x[ExtractRedundantConjunct-24]
	_ = x[CommuteVarInequality-25]
	_ = x[CommuteConstInequality-26]
	_ = x[NormalizeCmpPlusConst-27]
	_ = x[NormalizeCmpMinusConst-28]
	_ = x[NormalizeCmpConstMinus-29]
	_ = x[NormalizeTupleEquality-30]
	_ = x[FoldNullComparisonLeft-31]
	_ = x[FoldNullComparisonRight-32]
	_ = x[FoldIsNull-33]
	_ = x[FoldNonNullIsNull-34]
	_ = x[FoldNullTupleIsTupleNull-35]
	_ = x[FoldNonNullTupleIsTupleNull-36]
	_ = x[FoldIsNotNull-37]
	_ = x[FoldNonNullIsNotNull-38]
	_ = x[FoldNonNullTupleIsTupleNotNull-39]
	_ = x[FoldNullTupleIsTupleNotNull-40]
	_ = x[CommuteNullIs-41]
	_ = x[NormalizeCmpTimeZoneFunction-42]
	_ = x[NormalizeCmpTimeZoneFunctionTZ-43]
	_ = x[FoldEqZeroSTDistance-44]
	_ = x[FoldCmpSTDistanceLeft-45]
	_ = x[FoldCmpSTDistanceRight-46]
	_ = x[FoldCmpSTMaxDistanceLeft-47]
	_ = x[FoldCmpSTMaxDistanceRight-48]
	_ = x[FoldEqTrue-49]
	_ = x[FoldEqFalse-50]
	_ = x[FoldNeTrue-51]
	_ = x[FoldNeFalse-52]
	_ = x[DecorrelateJoin-53]
	_ = x[DecorrelateProjectSet-54]
	_ = x[TryDecorrelateSelect-55]
	_ = x[TryDecorrelateProject-56]
	_ = x[TryDecorrelateProjectSelect-57]
	_ = x[TryDecorrelateProjectInnerJoin-58]
	_ = x[TryDecorrelateInnerJoin-59]
	_ = x[TryDecorrelateInnerLeftJoin-60]
	_ = x[TryDecorrelateGroupBy-61]
	_ = x[TryDecorrelateScalarGroupBy-62]
	_ = x[TryDecorrelateSemiJoin-63]
	_ = x[TryDecorrelateLimitOne-64]
	_ = x[TryDecorrelateProjectSet-65]
	_ = x[TryDecorrelateWindow-66]
	_ = x[TryDecorrelateMax1Row-67]
	_ = x[HoistSelectExists-68]
	_ = x[HoistSelectNotExists-69]
	_ = x[HoistSelectSubquery-70]
	_ = x[HoistProjectSubquery-71]
	_ = x[HoistJoinSubquery-72]
	_ = x[HoistValuesSubquery-73]
	_ = x[HoistProjectSetSubquery-74]
	_ = x[NormalizeSelectAnyFilter-75]
	_ = x[NormalizeJoinAnyFilter-76]
	_ = x[NormalizeSelectNotAnyFilter-77]
	_ = x[NormalizeJoinNotAnyFilter-78]
	_ = x[FoldNullCast-79]
	_ = x[FoldNullUnary-80]
	_ = x[FoldNullBinaryLeft-81]
	_ = x[FoldNullBinaryRight-82]
	_ = x[FoldNullInNonEmpty-83]
	_ = x[FoldInEmpty-84]
	_ = x[FoldNotInEmpty-85]
	_ = x[FoldArray-86]
	_ = x[FoldBinary-87]
	_ = x[FoldUnary-88]
	_ = x[FoldComparison-89]
	_ = x[FoldCast-90]
	_ = x[FoldIndirection-91]
	_ = x[FoldColumnAccess-92]
	_ = x[FoldFunction-93]
	_ = x[FoldEqualsAnyNull-94]
	_ = x[ConvertGroupByToDistinct-95]
	_ = x[EliminateGroupByProject-96]
	_ = x[EliminateJoinUnderGroupByLeft-97]
	_ = x[EliminateJoinUnderGroupByRight-98]
	_ = x[EliminateDistinct-99]
	_ = x[ReduceGroupingCols-100]
	_ = x[ReduceNotNullGroupingCols-101]
	_ = x[EliminateAggDistinctForKeys-102]
	_ = x[EliminateAggFilteredDistinctForKeys-103]
	_ = x[EliminateDistinctNoColumns-104]
	_ = x[EliminateEnsureDistinctNoColumns-105]
	_ = x[EliminateDistinctOnValues-106]
	_ = x[PushAggDistinctIntoGroupBy-107]
	_ = x[PushAggFilterIntoScalarGroupBy-108]
	_ = x[ConvertCountToCountRows-109]
	_ = x[FoldGroupingOperators-110]
	_ = x[InlineConstVar-111]
	_ = x[InlineProjectConstants-112]
	_ = x[InlineSelectConstants-113]
	_ = x[InlineJoinConstantsLeft-114]
	_ = x[InlineJoinConstantsRight-115]
	_ = x[PushSelectIntoInlinableProject-116]
	_ = x[InlineProjectInProject-117]
	_ = x[CommuteRightJoin-118]
	_ = x[SimplifyJoinFilters-119]
	_ = x[DetectJoinContradiction-120]
	_ = x[PushFilterIntoJoinLeftAndRight-121]
	_ = x[MapFilterIntoJoinLeft-122]
	_ = x[MapFilterIntoJoinRight-123]
	_ = x[MapEqualityIntoJoinLeftAndRight-124]
	_ = x[PushFilterIntoJoinLeft-125]
	_ = x[PushFilterIntoJoinRight-126]
	_ = x[SimplifyLeftJoin-127]
	_ = x[SimplifyRightJoin-128]
	_ = x[EliminateSemiJoin-129]
	_ = x[SimplifyZeroCardinalitySemiJoin-130]
	_ = x[EliminateAntiJoin-131]
	_ = x[SimplifyZeroCardinalityAntiJoin-132]
	_ = x[EliminateJoinNoColsLeft-133]
	_ = x[EliminateJoinNoColsRight-134]
	_ = x[HoistJoinProjectRight-135]
	_ = x[HoistJoinProjectLeft-136]
	_ = x[SimplifyJoinNotNullEquality-137]
	_ = x[ExtractJoinEqualities-138]
	_ = x[SortFiltersInJoin-139]
	_ = x[LeftAssociateJoinsLeft-140]
	_ = x[LeftAssociateJoinsRight-141]
	_ = x[RightAssociateJoinsLeft-142]
	_ = x[RightAssociateJoinsRight-143]
	_ = x[RemoveJoinNotNullCondition-144]
	_ = x[EliminateLimit-145]
	_ = x[EliminateOffset-146]
	_ = x[PushLimitIntoProject-147]
	_ = x[PushOffsetIntoProject-148]
	_ = x[PushLimitIntoOffset-149]
	_ = x[PushLimitIntoOrdinality-150]
	_ = x[PushLimitIntoJoinLeft-151]
	_ = x[PushLimitIntoJoinRight-152]
	_ = x[FoldLimits-153]
	_ = x[AssociateLimitJoinsLeft-154]
	_ = x[AssociateLimitJoinsRight-155]
	_ = x[EliminateMax1Row-156]
	_ = x[FoldPlusZero-157]
	_ = x[FoldZeroPlus-158]
	_ = x[FoldMinusZero-159]
	_ = x[FoldMultOne-160]
	_ = x[FoldOneMult-161]
	_ = x[FoldDivOne-162]
	_ = x[InvertMinus-163]
	_ = x[EliminateUnaryMinus-164]
	_ = x[SimplifyLimitOrdering-165]
	_ = x[SimplifyOffsetOrdering-166]
	_ = x[SimplifyGroupByOrdering-167]
	_ = x[SimplifyOrdinalityOrdering-168]
	_ = x[SimplifyExplainOrdering-169]
	_ = x[EliminateJoinUnderProjectLeft-170]
	_ = x[EliminateJoinUnderProjectRight-171]
	_ = x[EliminateProject-172]
	_ = x[MergeProjects-173]
	_ = x[MergeProjectWithValues-174]
	_ = x[PushColumnRemappingIntoValues-175]
	_ = x[FoldTupleAccessIntoValues-176]
	_ = x[FoldJSONAccessIntoValues-177]
	_ = x[ConvertZipArraysToValues-178]
	_ = x[PruneProjectCols-179]
	_ = x[PruneScanCols-180]
	_ = x[PruneSelectCols-181]
	_ = x[PruneLimitCols-182]
	_ = x[PruneOffsetCols-183]
	_ = x[PruneJoinLeftCols-184]
	_ = x[PruneJoinRightCols-185]
	_ = x[PruneSemiAntiJoinRightCols-186]
	_ = x[PruneAggCols-187]
	_ = x[PruneGroupByCols-188]
	_ = x[PruneValuesCols-189]
	_ = x[PruneOrdinalityCols-190]
	_ = x[PruneExplainCols-191]
	_ = x[PruneProjectSetCols-192]
	_ = x[PruneWindowOutputCols-193]
	_ = x[PruneWindowInputCols-194]
	_ = x[PruneMutationFetchCols-195]
	_ = x[PruneMutationInputCols-196]
	_ = x[PruneMutationReturnCols-197]
	_ = x[PruneWithScanCols-198]
	_ = x[PruneWithCols-199]
	_ = x[PruneUnionAllCols-200]
	_ = x[RejectNullsLeftJoin-201]
	_ = x[RejectNullsRightJoin-202]
	_ = x[RejectNullsGroupBy-203]
	_ = x[RejectNullsUnderJoinLeft-204]
	_ = x[RejectNullsUnderJoinRight-205]
	_ = x[RejectNullsProject-206]
	_ = x[CommuteVar-207]
	_ = x[CommuteConst-208]
	_ = x[EliminateCoalesce-209]
	_ = x[SimplifyCoalesce-210]
	_ = x[EliminateCast-211]
	_ = x[NormalizeInConst-212]
	_ = x[FoldInNull-213]
	_ = x[UnifyComparisonTypes-214]
	_ = x[EliminateExistsZeroRows-215]
	_ = x[EliminateExistsProject-216]
	_ = x[EliminateExistsGroupBy-217]
	_ = x[InlineExistsSelectTuple-218]
	_ = x[IntroduceExistsLimit-219]
	_ = x[EliminateExistsLimit-220]
	_ = x[SimplifyCaseWhenConstValue-221]
	_ = x[InlineAnyValuesSingleCol-222]
	_ = x[InlineAnyValuesMultiCol-223]
	_ = x[SimplifyEqualsAnyTuple-224]
	_ = x[SimplifyAnyScalarArray-225]
	_ = x[FoldCollate-226]
	_ = x[NormalizeArrayFlattenToAgg-227]
	_ = x[SimplifySameVarEqualities-228]
	_ = x[SimplifySameVarInequalities-229]
	_ = x[SimplifySelectFilters-230]
	_ = x[ConsolidateSelectFilters-231]
	_ = x[DetectSelectContradiction-232]
	_ = x[EliminateSelect-233]
	_ = x[MergeSelects-234]
	_ = x[PushSelectIntoProject-235]
	_ = x[MergeSelectInnerJoin-236]
	_ = x[PushSelectCondLeftIntoJoinLeftAndRight-237]
	_ = x[PushSelectIntoJoinLeft-238]
	_ = x[PushSelectIntoGroupBy-239]
	_ = x[RemoveNotNullCondition-240]
	_ = x[PushSelectIntoProjectSet-241]
	_ = x[PushFilterIntoSetOp-242]
	_ = x[EliminateUnionAllLeft-243]
	_ = x[EliminateUnionAllRight-244]
	_ = x[EliminateWindow-245]
	_ = x[ReduceWindowPartitionCols-246]
	_ = x[SimplifyWindowOrdering-247]
	_ = x[PushSelectIntoWindow-248]
	_ = x[PushLimitIntoWindow-249]
	_ = x[InlineWith-250]
	_ = x[startExploreRule-251]
	_ = x[ReplaceScalarMinMaxWithLimit-252]
	_ = x[ReplaceMinWithLimit-253]
	_ = x[ReplaceMaxWithLimit-254]
	_ = x[GenerateStreamingGroupBy-255]
	_ = x[ReorderJoins-256]
	_ = x[CommuteLeftJoin-257]
	_ = x[CommuteSemiJoin-258]
	_ = x[ConvertSemiToInnerJoin-259]
	_ = x[ConvertLeftToInnerJoin-260]
	_ = x[ConvertAntiToLeftJoin-261]
	_ = x[GenerateMergeJoins-262]
	_ = x[GenerateLookupJoins-263]
	_ = x[GenerateInvertedJoins-264]
	_ = x[GenerateInvertedJoinsFromSelect-265]
	_ = x[GenerateZigzagJoins-266]
	_ = x[GenerateInvertedIndexZigzagJoins-267]
	_ = x[GenerateLookupJoinsWithFilter-268]
	_ = x[PushJoinIntoIndexJoin-269]
	_ = x[GenerateLimitedScans-270]
	_ = x[PushLimitIntoFilteredScan-271]
	_ = x[PushLimitIntoIndexJoin-272]
	_ = x[SplitScanIntoUnionScans-273]
	_ = x[EliminateIndexJoinInsideProject-274]
	_ = x[GenerateIndexScans-275]
	_ = x[GeneratePartialIndexScans-276]
	_ = x[GenerateConstrainedScans-277]
	_ = x[GenerateInvertedIndexScans-278]
	_ = x[SplitDisjunction-279]
	_ = x[SplitDisjunctionAddKey-280]
	_ = x[NumRuleNames-281]
}

const _RuleName_name = "InvalidRuleNameSimplifyRootOrderingPruneRootColsSimplifyZeroCardinalityGroupNumManualRuleNamesEliminateAggDistinctNormalizeNestedAndsSimplifyTrueAndSimplifyAndTrueSimplifyFalseAndSimplifyAndFalseSimplifyTrueOrSimplifyOrTrueSimplifyFalseOrSimplifyOrFalseSimplifyRangeFoldNullAndOrFoldNotTrueFoldNotFalseFoldNotNullNegateComparisonEliminateNotNegateAndNegateOrExtractRedundantConjunctCommuteVarInequalityCommuteConstInequalityNormalizeCmpPlusConstNormalizeCmpMinusConstNormalizeCmpConstMinusNormalizeTupleEqualityFoldNullComparisonLeftFoldNullComparisonRightFoldIsNullFoldNonNullIsNullFoldNullTupleIsTupleNullFoldNonNullTupleIsTupleNullFoldIsNotNullFoldNonNullIsNotNullFoldNonNullTupleIsTupleNotNullFoldNullTupleIsTupleNotNullCommuteNullIsNormalizeCmpTimeZoneFunctionNormalizeCmpTimeZoneFunctionTZFoldEqZeroSTDistanceFoldCmpSTDistanceLeftFoldCmpSTDistanceRightFoldCmpSTMaxDistanceLeftFoldCmpSTMaxDistanceRightFoldEqTrueFoldEqFalseFoldNeTrueFoldNeFalseDecorrelateJoinDecorrelateProjectSetTryDecorrelateSelectTryDecorrelateProjectTryDecorrelateProjectSelectTryDecorrelateProjectInnerJoinTryDecorrelateInnerJoinTryDecorrelateInnerLeftJoinTryDecorrelateGroupByTryDecorrelateScalarGroupByTryDecorrelateSemiJoinTryDecorrelateLimitOneTryDecorrelateProjectSetTryDecorrelateWindowTryDecorrelateMax1RowHoistSelectExistsHoistSelectNotExistsHoistSelectSubqueryHoistProjectSubqueryHoistJoinSubqueryHoistValuesSubqueryHoistProjectSetSubqueryNormalizeSelectAnyFilterNormalizeJoinAnyFilterNormalizeSelectNotAnyFilterNormalizeJoinNotAnyFilterFoldNullCastFoldNullUnaryFoldNullBinaryLeftFoldNullBinaryRightFoldNullInNonEmptyFoldInEmptyFoldNotInEmptyFoldArrayFoldBinaryFoldUnaryFoldComparisonFoldCastFoldIndirectionFoldColumnAccessFoldFunctionFoldEqualsAnyNullConvertGroupByToDistinctEliminateGroupByProjectEliminateJoinUnderGroupByLeftEliminateJoinUnderGroupByRightEliminateDistinctReduceGroupingColsReduceNotNullGroupingColsEliminateAggDistinctForKeysEliminateAggFilteredDistinctForKeysEliminateDistinctNoColumnsEliminateEnsureDistinctNoColumnsEliminateDistinctOnValuesPushAggDistinctIntoGroupByPushAggFilterIntoScalarGroupByConvertCountToCountRowsFoldGroupingOperatorsInlineConstVarInlineProjectConstantsInlineSelectConstantsInlineJoinConstantsLeftInlineJoinConstantsRightPushSelectIntoInlinableProjectInlineProjectInProjectCommuteRightJoinSimplifyJoinFiltersDetectJoinContradictionPushFilterIntoJoinLeftAndRightMapFilterIntoJoinLeftMapFilterIntoJoinRightMapEqualityIntoJoinLeftAndRightPushFilterIntoJoinLeftPushFilterIntoJoinRightSimplifyLeftJoinSimplifyRightJoinEliminateSemiJoinSimplifyZeroCardinalitySemiJoinEliminateAntiJoinSimplifyZeroCardinalityAntiJoinEliminateJoinNoColsLeftEliminateJoinNoColsRightHoistJoinProjectRightHoistJoinProjectLeftSimplifyJoinNotNullEqualityExtractJoinEqualitiesSortFiltersInJoinLeftAssociateJoinsLeftLeftAssociateJoinsRightRightAssociateJoinsLeftRightAssociateJoinsRightRemoveJoinNotNullConditionEliminateLimitEliminateOffsetPushLimitIntoProjectPushOffsetIntoProjectPushLimitIntoOffsetPushLimitIntoOrdinalityPushLimitIntoJoinLeftPushLimitIntoJoinRightFoldLimitsAssociateLimitJoinsLeftAssociateLimitJoinsRightEliminateMax1RowFoldPlusZeroFoldZeroPlusFoldMinusZeroFoldMultOneFoldOneMultFoldDivOneInvertMinusEliminateUnaryMinusSimplifyLimitOrderingSimplifyOffsetOrderingSimplifyGroupByOrderingSimplifyOrdinalityOrderingSimplifyExplainOrderingEliminateJoinUnderProjectLeftEliminateJoinUnderProjectRightEliminateProjectMergeProjectsMergeProjectWithValuesPushColumnRemappingIntoValuesFoldTupleAccessIntoValuesFoldJSONAccessIntoValuesConvertZipArraysToValuesPruneProjectColsPruneScanColsPruneSelectColsPruneLimitColsPruneOffsetColsPruneJoinLeftColsPruneJoinRightColsPruneSemiAntiJoinRightColsPruneAggColsPruneGroupByColsPruneValuesColsPruneOrdinalityColsPruneExplainColsPruneProjectSetColsPruneWindowOutputColsPruneWindowInputColsPruneMutationFetchColsPruneMutationInputColsPruneMutationReturnColsPruneWithScanColsPruneWithColsPruneUnionAllColsRejectNullsLeftJoinRejectNullsRightJoinRejectNullsGroupByRejectNullsUnderJoinLeftRejectNullsUnderJoinRightRejectNullsProjectCommuteVarCommuteConstEliminateCoalesceSimplifyCoalesceEliminateCastNormalizeInConstFoldInNullUnifyComparisonTypesEliminateExistsZeroRowsEliminateExistsProjectEliminateExistsGroupByInlineExistsSelectTupleIntroduceExistsLimitEliminateExistsLimitSimplifyCaseWhenConstValueInlineAnyValuesSingleColInlineAnyValuesMultiColSimplifyEqualsAnyTupleSimplifyAnyScalarArrayFoldCollateNormalizeArrayFlattenToAggSimplifySameVarEqualitiesSimplifySameVarInequalitiesSimplifySelectFiltersConsolidateSelectFiltersDetectSelectContradictionEliminateSelectMergeSelectsPushSelectIntoProjectMergeSelectInnerJoinPushSelectCondLeftIntoJoinLeftAndRightPushSelectIntoJoinLeftPushSelectIntoGroupByRemoveNotNullConditionPushSelectIntoProjectSetPushFilterIntoSetOpEliminateUnionAllLeftEliminateUnionAllRightEliminateWindowReduceWindowPartitionColsSimplifyWindowOrderingPushSelectIntoWindowPushLimitIntoWindowInlineWithstartExploreRuleReplaceScalarMinMaxWithLimitReplaceMinWithLimitReplaceMaxWithLimitGenerateStreamingGroupByReorderJoinsCommuteLeftJoinCommuteSemiJoinConvertSemiToInnerJoinConvertLeftToInnerJoinConvertAntiToLeftJoinGenerateMergeJoinsGenerateLookupJoinsGenerateInvertedJoinsGenerateInvertedJoinsFromSelectGenerateZigzagJoinsGenerateInvertedIndexZigzagJoinsGenerateLookupJoinsWithFilterPushJoinIntoIndexJoinGenerateLimitedScansPushLimitIntoFilteredScanPushLimitIntoIndexJoinSplitScanIntoUnionScansEliminateIndexJoinInsideProjectGenerateIndexScansGeneratePartialIndexScansGenerateConstrainedScansGenerateInvertedIndexScansSplitDisjunctionSplitDisjunctionAddKeyNumRuleNames"

var _RuleName_index = [...]uint16{0, 15, 35, 48, 76, 94, 114, 133, 148, 163, 179, 195, 209, 223, 238, 253, 266, 279, 290, 302, 313, 329, 341, 350, 358, 382, 402, 424, 445, 467, 489, 511, 533, 556, 566, 583, 607, 634, 647, 667, 697, 724, 737, 765, 795, 815, 836, 858, 882, 907, 917, 928, 938, 949, 964, 985, 1005, 1026, 1053, 1083, 1106, 1133, 1154, 1181, 1203, 1225, 1249, 1269, 1290, 1307, 1327, 1346, 1366, 1383, 1402, 1425, 1449, 1471, 1498, 1523, 1535, 1548, 1566, 1585, 1603, 1614, 1628, 1637, 1647, 1656, 1670, 1678, 1693, 1709, 1721, 1738, 1762, 1785, 1814, 1844, 1861, 1879, 1904, 1931, 1966, 1992, 2024, 2049, 2075, 2105, 2128, 2149, 2163, 2185, 2206, 2229, 2253, 2283, 2305, 2321, 2340, 2363, 2393, 2414, 2436, 2467, 2489, 2512, 2528, 2545, 2562, 2593, 2610, 2641, 2664, 2688, 2709, 2729, 2756, 2777, 2794, 2816, 2839, 2862, 2886, 2912, 2926, 2941, 2961, 2982, 3001, 3024, 3045, 3067, 3077, 3100, 3124, 3140, 3152, 3164, 3177, 3188, 3199, 3209, 3220, 3239, 3260, 3282, 3305, 3331, 3354, 3383, 3413, 3429, 3442, 3464, 3493, 3518, 3542, 3566, 3582, 3595, 3610, 3624, 3639, 3656, 3674, 3700, 3712, 3728, 3743, 3762, 3778, 3797, 3818, 3838, 3860, 3882, 3905, 3922, 3935, 3952, 3971, 3991, 4009, 4033, 4058, 4076, 4086, 4098, 4115, 4131, 4144, 4160, 4170, 4190, 4213, 4235, 4257, 4280, 4300, 4320, 4346, 4370, 4393, 4415, 4437, 4448, 4474, 4499, 4526, 4547, 4571, 4596, 4611, 4623, 4644, 4664, 4702, 4724, 4745, 4767, 4791, 4810, 4831, 4853, 4868, 4893, 4915, 4935, 4954, 4964, 4980, 5008, 5027, 5046, 5070, 5082, 5097, 5112, 5134, 5156, 5177, 5195, 5214, 5235, 5266, 5285, 5317, 5346, 5367, 5387, 5412, 5434, 5457, 5488, 5506, 5531, 5555, 5581, 5597, 5619, 5631}

func (i RuleName) String() string {
	if i >= RuleName(len(_RuleName_index)-1) {
		return "RuleName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RuleName_name[_RuleName_index[i]:_RuleName_index[i+1]]
}
