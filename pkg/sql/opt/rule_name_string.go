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
	_ = x[DecorrelateJoin-49]
	_ = x[DecorrelateProjectSet-50]
	_ = x[TryDecorrelateSelect-51]
	_ = x[TryDecorrelateProject-52]
	_ = x[TryDecorrelateProjectSelect-53]
	_ = x[TryDecorrelateProjectInnerJoin-54]
	_ = x[TryDecorrelateInnerJoin-55]
	_ = x[TryDecorrelateInnerLeftJoin-56]
	_ = x[TryDecorrelateGroupBy-57]
	_ = x[TryDecorrelateScalarGroupBy-58]
	_ = x[TryDecorrelateSemiJoin-59]
	_ = x[TryDecorrelateLimitOne-60]
	_ = x[TryDecorrelateProjectSet-61]
	_ = x[TryDecorrelateWindow-62]
	_ = x[TryDecorrelateMax1Row-63]
	_ = x[HoistSelectExists-64]
	_ = x[HoistSelectNotExists-65]
	_ = x[HoistSelectSubquery-66]
	_ = x[HoistProjectSubquery-67]
	_ = x[HoistJoinSubquery-68]
	_ = x[HoistValuesSubquery-69]
	_ = x[HoistProjectSetSubquery-70]
	_ = x[NormalizeSelectAnyFilter-71]
	_ = x[NormalizeJoinAnyFilter-72]
	_ = x[NormalizeSelectNotAnyFilter-73]
	_ = x[NormalizeJoinNotAnyFilter-74]
	_ = x[FoldNullCast-75]
	_ = x[FoldNullUnary-76]
	_ = x[FoldNullBinaryLeft-77]
	_ = x[FoldNullBinaryRight-78]
	_ = x[FoldNullInNonEmpty-79]
	_ = x[FoldInEmpty-80]
	_ = x[FoldNotInEmpty-81]
	_ = x[FoldArray-82]
	_ = x[FoldBinary-83]
	_ = x[FoldUnary-84]
	_ = x[FoldComparison-85]
	_ = x[FoldCast-86]
	_ = x[FoldIndirection-87]
	_ = x[FoldColumnAccess-88]
	_ = x[FoldFunction-89]
	_ = x[FoldEqualsAnyNull-90]
	_ = x[ConvertGroupByToDistinct-91]
	_ = x[EliminateGroupByProject-92]
	_ = x[EliminateJoinUnderGroupByLeft-93]
	_ = x[EliminateJoinUnderGroupByRight-94]
	_ = x[EliminateDistinct-95]
	_ = x[ReduceGroupingCols-96]
	_ = x[ReduceNotNullGroupingCols-97]
	_ = x[EliminateAggDistinctForKeys-98]
	_ = x[EliminateAggFilteredDistinctForKeys-99]
	_ = x[EliminateDistinctNoColumns-100]
	_ = x[EliminateEnsureDistinctNoColumns-101]
	_ = x[EliminateDistinctOnValues-102]
	_ = x[PushAggDistinctIntoGroupBy-103]
	_ = x[PushAggFilterIntoScalarGroupBy-104]
	_ = x[ConvertCountToCountRows-105]
	_ = x[FoldGroupingOperators-106]
	_ = x[InlineConstVar-107]
	_ = x[InlineProjectConstants-108]
	_ = x[InlineSelectConstants-109]
	_ = x[InlineJoinConstantsLeft-110]
	_ = x[InlineJoinConstantsRight-111]
	_ = x[PushSelectIntoInlinableProject-112]
	_ = x[InlineProjectInProject-113]
	_ = x[CommuteRightJoin-114]
	_ = x[SimplifyJoinFilters-115]
	_ = x[DetectJoinContradiction-116]
	_ = x[PushFilterIntoJoinLeftAndRight-117]
	_ = x[MapFilterIntoJoinLeft-118]
	_ = x[MapFilterIntoJoinRight-119]
	_ = x[MapEqualityIntoJoinLeftAndRight-120]
	_ = x[PushFilterIntoJoinLeft-121]
	_ = x[PushFilterIntoJoinRight-122]
	_ = x[SimplifyLeftJoin-123]
	_ = x[SimplifyRightJoin-124]
	_ = x[EliminateSemiJoin-125]
	_ = x[SimplifyZeroCardinalitySemiJoin-126]
	_ = x[EliminateAntiJoin-127]
	_ = x[SimplifyZeroCardinalityAntiJoin-128]
	_ = x[EliminateJoinNoColsLeft-129]
	_ = x[EliminateJoinNoColsRight-130]
	_ = x[HoistJoinProjectRight-131]
	_ = x[HoistJoinProjectLeft-132]
	_ = x[SimplifyJoinNotNullEquality-133]
	_ = x[ExtractJoinEqualities-134]
	_ = x[SortFiltersInJoin-135]
	_ = x[LeftAssociateJoinsLeft-136]
	_ = x[LeftAssociateJoinsRight-137]
	_ = x[RightAssociateJoinsLeft-138]
	_ = x[RightAssociateJoinsRight-139]
	_ = x[RemoveJoinNotNullCondition-140]
	_ = x[EliminateLimit-141]
	_ = x[EliminateOffset-142]
	_ = x[PushLimitIntoProject-143]
	_ = x[PushOffsetIntoProject-144]
	_ = x[PushLimitIntoOffset-145]
	_ = x[PushLimitIntoOrdinality-146]
	_ = x[PushLimitIntoJoinLeft-147]
	_ = x[PushLimitIntoJoinRight-148]
	_ = x[FoldLimits-149]
	_ = x[AssociateLimitJoinsLeft-150]
	_ = x[AssociateLimitJoinsRight-151]
	_ = x[EliminateMax1Row-152]
	_ = x[FoldPlusZero-153]
	_ = x[FoldZeroPlus-154]
	_ = x[FoldMinusZero-155]
	_ = x[FoldMultOne-156]
	_ = x[FoldOneMult-157]
	_ = x[FoldDivOne-158]
	_ = x[InvertMinus-159]
	_ = x[EliminateUnaryMinus-160]
	_ = x[SimplifyLimitOrdering-161]
	_ = x[SimplifyOffsetOrdering-162]
	_ = x[SimplifyGroupByOrdering-163]
	_ = x[SimplifyOrdinalityOrdering-164]
	_ = x[SimplifyExplainOrdering-165]
	_ = x[EliminateJoinUnderProjectLeft-166]
	_ = x[EliminateJoinUnderProjectRight-167]
	_ = x[EliminateProject-168]
	_ = x[MergeProjects-169]
	_ = x[MergeProjectWithValues-170]
	_ = x[PushColumnRemappingIntoValues-171]
	_ = x[FoldTupleAccessIntoValues-172]
	_ = x[FoldJSONAccessIntoValues-173]
	_ = x[ConvertZipArraysToValues-174]
	_ = x[PruneProjectCols-175]
	_ = x[PruneScanCols-176]
	_ = x[PruneSelectCols-177]
	_ = x[PruneLimitCols-178]
	_ = x[PruneOffsetCols-179]
	_ = x[PruneJoinLeftCols-180]
	_ = x[PruneJoinRightCols-181]
	_ = x[PruneSemiAntiJoinRightCols-182]
	_ = x[PruneAggCols-183]
	_ = x[PruneGroupByCols-184]
	_ = x[PruneValuesCols-185]
	_ = x[PruneOrdinalityCols-186]
	_ = x[PruneExplainCols-187]
	_ = x[PruneProjectSetCols-188]
	_ = x[PruneWindowOutputCols-189]
	_ = x[PruneWindowInputCols-190]
	_ = x[PruneMutationFetchCols-191]
	_ = x[PruneMutationInputCols-192]
	_ = x[PruneMutationReturnCols-193]
	_ = x[PruneWithScanCols-194]
	_ = x[PruneWithCols-195]
	_ = x[PruneUnionAllCols-196]
	_ = x[RejectNullsLeftJoin-197]
	_ = x[RejectNullsRightJoin-198]
	_ = x[RejectNullsGroupBy-199]
	_ = x[RejectNullsUnderJoinLeft-200]
	_ = x[RejectNullsUnderJoinRight-201]
	_ = x[RejectNullsProject-202]
	_ = x[CommuteVar-203]
	_ = x[CommuteConst-204]
	_ = x[EliminateCoalesce-205]
	_ = x[SimplifyCoalesce-206]
	_ = x[EliminateCast-207]
	_ = x[NormalizeInConst-208]
	_ = x[FoldInNull-209]
	_ = x[UnifyComparisonTypes-210]
	_ = x[EliminateExistsZeroRows-211]
	_ = x[EliminateExistsProject-212]
	_ = x[EliminateExistsGroupBy-213]
	_ = x[InlineExistsSelectTuple-214]
	_ = x[IntroduceExistsLimit-215]
	_ = x[EliminateExistsLimit-216]
	_ = x[SimplifyCaseWhenConstValue-217]
	_ = x[InlineAnyValuesSingleCol-218]
	_ = x[InlineAnyValuesMultiCol-219]
	_ = x[SimplifyEqualsAnyTuple-220]
	_ = x[SimplifyAnyScalarArray-221]
	_ = x[FoldCollate-222]
	_ = x[NormalizeArrayFlattenToAgg-223]
	_ = x[SimplifySameVarEqualities-224]
	_ = x[SimplifySameVarInequalities-225]
	_ = x[SimplifySelectFilters-226]
	_ = x[ConsolidateSelectFilters-227]
	_ = x[DetectSelectContradiction-228]
	_ = x[EliminateSelect-229]
	_ = x[MergeSelects-230]
	_ = x[PushSelectIntoProject-231]
	_ = x[MergeSelectInnerJoin-232]
	_ = x[PushSelectCondLeftIntoJoinLeftAndRight-233]
	_ = x[PushSelectIntoJoinLeft-234]
	_ = x[PushSelectIntoGroupBy-235]
	_ = x[RemoveNotNullCondition-236]
	_ = x[PushSelectIntoProjectSet-237]
	_ = x[PushFilterIntoSetOp-238]
	_ = x[EliminateUnionAllLeft-239]
	_ = x[EliminateUnionAllRight-240]
	_ = x[EliminateWindow-241]
	_ = x[ReduceWindowPartitionCols-242]
	_ = x[SimplifyWindowOrdering-243]
	_ = x[PushSelectIntoWindow-244]
	_ = x[PushLimitIntoWindow-245]
	_ = x[InlineWith-246]
	_ = x[startExploreRule-247]
	_ = x[ReplaceScalarMinMaxWithLimit-248]
	_ = x[ReplaceMinWithLimit-249]
	_ = x[ReplaceMaxWithLimit-250]
	_ = x[GenerateStreamingGroupBy-251]
	_ = x[ReorderJoins-252]
	_ = x[CommuteLeftJoin-253]
	_ = x[CommuteSemiJoin-254]
	_ = x[ConvertSemiToInnerJoin-255]
	_ = x[ConvertLeftToInnerJoin-256]
	_ = x[ConvertAntiToLeftJoin-257]
	_ = x[GenerateMergeJoins-258]
	_ = x[GenerateLookupJoins-259]
	_ = x[GenerateInvertedJoins-260]
	_ = x[GenerateInvertedJoinsFromSelect-261]
	_ = x[GenerateZigzagJoins-262]
	_ = x[GenerateInvertedIndexZigzagJoins-263]
	_ = x[GenerateLookupJoinsWithFilter-264]
	_ = x[PushJoinIntoIndexJoin-265]
	_ = x[GenerateLimitedScans-266]
	_ = x[PushLimitIntoFilteredScan-267]
	_ = x[PushLimitIntoIndexJoin-268]
	_ = x[SplitScanIntoUnionScans-269]
	_ = x[EliminateIndexJoinInsideProject-270]
	_ = x[GenerateIndexScans-271]
	_ = x[GeneratePartialIndexScans-272]
	_ = x[GenerateConstrainedScans-273]
	_ = x[GenerateInvertedIndexScans-274]
	_ = x[SplitDisjunction-275]
	_ = x[SplitDisjunctionAddKey-276]
	_ = x[NumRuleNames-277]
}

const _RuleName_name = "InvalidRuleNameSimplifyRootOrderingPruneRootColsSimplifyZeroCardinalityGroupNumManualRuleNamesEliminateAggDistinctNormalizeNestedAndsSimplifyTrueAndSimplifyAndTrueSimplifyFalseAndSimplifyAndFalseSimplifyTrueOrSimplifyOrTrueSimplifyFalseOrSimplifyOrFalseSimplifyRangeFoldNullAndOrFoldNotTrueFoldNotFalseFoldNotNullNegateComparisonEliminateNotNegateAndNegateOrExtractRedundantConjunctCommuteVarInequalityCommuteConstInequalityNormalizeCmpPlusConstNormalizeCmpMinusConstNormalizeCmpConstMinusNormalizeTupleEqualityFoldNullComparisonLeftFoldNullComparisonRightFoldIsNullFoldNonNullIsNullFoldNullTupleIsTupleNullFoldNonNullTupleIsTupleNullFoldIsNotNullFoldNonNullIsNotNullFoldNonNullTupleIsTupleNotNullFoldNullTupleIsTupleNotNullCommuteNullIsNormalizeCmpTimeZoneFunctionNormalizeCmpTimeZoneFunctionTZFoldEqZeroSTDistanceFoldCmpSTDistanceLeftFoldCmpSTDistanceRightFoldCmpSTMaxDistanceLeftFoldCmpSTMaxDistanceRightDecorrelateJoinDecorrelateProjectSetTryDecorrelateSelectTryDecorrelateProjectTryDecorrelateProjectSelectTryDecorrelateProjectInnerJoinTryDecorrelateInnerJoinTryDecorrelateInnerLeftJoinTryDecorrelateGroupByTryDecorrelateScalarGroupByTryDecorrelateSemiJoinTryDecorrelateLimitOneTryDecorrelateProjectSetTryDecorrelateWindowTryDecorrelateMax1RowHoistSelectExistsHoistSelectNotExistsHoistSelectSubqueryHoistProjectSubqueryHoistJoinSubqueryHoistValuesSubqueryHoistProjectSetSubqueryNormalizeSelectAnyFilterNormalizeJoinAnyFilterNormalizeSelectNotAnyFilterNormalizeJoinNotAnyFilterFoldNullCastFoldNullUnaryFoldNullBinaryLeftFoldNullBinaryRightFoldNullInNonEmptyFoldInEmptyFoldNotInEmptyFoldArrayFoldBinaryFoldUnaryFoldComparisonFoldCastFoldIndirectionFoldColumnAccessFoldFunctionFoldEqualsAnyNullConvertGroupByToDistinctEliminateGroupByProjectEliminateJoinUnderGroupByLeftEliminateJoinUnderGroupByRightEliminateDistinctReduceGroupingColsReduceNotNullGroupingColsEliminateAggDistinctForKeysEliminateAggFilteredDistinctForKeysEliminateDistinctNoColumnsEliminateEnsureDistinctNoColumnsEliminateDistinctOnValuesPushAggDistinctIntoGroupByPushAggFilterIntoScalarGroupByConvertCountToCountRowsFoldGroupingOperatorsInlineConstVarInlineProjectConstantsInlineSelectConstantsInlineJoinConstantsLeftInlineJoinConstantsRightPushSelectIntoInlinableProjectInlineProjectInProjectCommuteRightJoinSimplifyJoinFiltersDetectJoinContradictionPushFilterIntoJoinLeftAndRightMapFilterIntoJoinLeftMapFilterIntoJoinRightMapEqualityIntoJoinLeftAndRightPushFilterIntoJoinLeftPushFilterIntoJoinRightSimplifyLeftJoinSimplifyRightJoinEliminateSemiJoinSimplifyZeroCardinalitySemiJoinEliminateAntiJoinSimplifyZeroCardinalityAntiJoinEliminateJoinNoColsLeftEliminateJoinNoColsRightHoistJoinProjectRightHoistJoinProjectLeftSimplifyJoinNotNullEqualityExtractJoinEqualitiesSortFiltersInJoinLeftAssociateJoinsLeftLeftAssociateJoinsRightRightAssociateJoinsLeftRightAssociateJoinsRightRemoveJoinNotNullConditionEliminateLimitEliminateOffsetPushLimitIntoProjectPushOffsetIntoProjectPushLimitIntoOffsetPushLimitIntoOrdinalityPushLimitIntoJoinLeftPushLimitIntoJoinRightFoldLimitsAssociateLimitJoinsLeftAssociateLimitJoinsRightEliminateMax1RowFoldPlusZeroFoldZeroPlusFoldMinusZeroFoldMultOneFoldOneMultFoldDivOneInvertMinusEliminateUnaryMinusSimplifyLimitOrderingSimplifyOffsetOrderingSimplifyGroupByOrderingSimplifyOrdinalityOrderingSimplifyExplainOrderingEliminateJoinUnderProjectLeftEliminateJoinUnderProjectRightEliminateProjectMergeProjectsMergeProjectWithValuesPushColumnRemappingIntoValuesFoldTupleAccessIntoValuesFoldJSONAccessIntoValuesConvertZipArraysToValuesPruneProjectColsPruneScanColsPruneSelectColsPruneLimitColsPruneOffsetColsPruneJoinLeftColsPruneJoinRightColsPruneSemiAntiJoinRightColsPruneAggColsPruneGroupByColsPruneValuesColsPruneOrdinalityColsPruneExplainColsPruneProjectSetColsPruneWindowOutputColsPruneWindowInputColsPruneMutationFetchColsPruneMutationInputColsPruneMutationReturnColsPruneWithScanColsPruneWithColsPruneUnionAllColsRejectNullsLeftJoinRejectNullsRightJoinRejectNullsGroupByRejectNullsUnderJoinLeftRejectNullsUnderJoinRightRejectNullsProjectCommuteVarCommuteConstEliminateCoalesceSimplifyCoalesceEliminateCastNormalizeInConstFoldInNullUnifyComparisonTypesEliminateExistsZeroRowsEliminateExistsProjectEliminateExistsGroupByInlineExistsSelectTupleIntroduceExistsLimitEliminateExistsLimitSimplifyCaseWhenConstValueInlineAnyValuesSingleColInlineAnyValuesMultiColSimplifyEqualsAnyTupleSimplifyAnyScalarArrayFoldCollateNormalizeArrayFlattenToAggSimplifySameVarEqualitiesSimplifySameVarInequalitiesSimplifySelectFiltersConsolidateSelectFiltersDetectSelectContradictionEliminateSelectMergeSelectsPushSelectIntoProjectMergeSelectInnerJoinPushSelectCondLeftIntoJoinLeftAndRightPushSelectIntoJoinLeftPushSelectIntoGroupByRemoveNotNullConditionPushSelectIntoProjectSetPushFilterIntoSetOpEliminateUnionAllLeftEliminateUnionAllRightEliminateWindowReduceWindowPartitionColsSimplifyWindowOrderingPushSelectIntoWindowPushLimitIntoWindowInlineWithstartExploreRuleReplaceScalarMinMaxWithLimitReplaceMinWithLimitReplaceMaxWithLimitGenerateStreamingGroupByReorderJoinsCommuteLeftJoinCommuteSemiJoinConvertSemiToInnerJoinConvertLeftToInnerJoinConvertAntiToLeftJoinGenerateMergeJoinsGenerateLookupJoinsGenerateInvertedJoinsGenerateInvertedJoinsFromSelectGenerateZigzagJoinsGenerateInvertedIndexZigzagJoinsGenerateLookupJoinsWithFilterPushJoinIntoIndexJoinGenerateLimitedScansPushLimitIntoFilteredScanPushLimitIntoIndexJoinSplitScanIntoUnionScansEliminateIndexJoinInsideProjectGenerateIndexScansGeneratePartialIndexScansGenerateConstrainedScansGenerateInvertedIndexScansSplitDisjunctionSplitDisjunctionAddKeyNumRuleNames"

var _RuleName_index = [...]uint16{0, 15, 35, 48, 76, 94, 114, 133, 148, 163, 179, 195, 209, 223, 238, 253, 266, 279, 290, 302, 313, 329, 341, 350, 358, 382, 402, 424, 445, 467, 489, 511, 533, 556, 566, 583, 607, 634, 647, 667, 697, 724, 737, 765, 795, 815, 836, 858, 882, 907, 922, 943, 963, 984, 1011, 1041, 1064, 1091, 1112, 1139, 1161, 1183, 1207, 1227, 1248, 1265, 1285, 1304, 1324, 1341, 1360, 1383, 1407, 1429, 1456, 1481, 1493, 1506, 1524, 1543, 1561, 1572, 1586, 1595, 1605, 1614, 1628, 1636, 1651, 1667, 1679, 1696, 1720, 1743, 1772, 1802, 1819, 1837, 1862, 1889, 1924, 1950, 1982, 2007, 2033, 2063, 2086, 2107, 2121, 2143, 2164, 2187, 2211, 2241, 2263, 2279, 2298, 2321, 2351, 2372, 2394, 2425, 2447, 2470, 2486, 2503, 2520, 2551, 2568, 2599, 2622, 2646, 2667, 2687, 2714, 2735, 2752, 2774, 2797, 2820, 2844, 2870, 2884, 2899, 2919, 2940, 2959, 2982, 3003, 3025, 3035, 3058, 3082, 3098, 3110, 3122, 3135, 3146, 3157, 3167, 3178, 3197, 3218, 3240, 3263, 3289, 3312, 3341, 3371, 3387, 3400, 3422, 3451, 3476, 3500, 3524, 3540, 3553, 3568, 3582, 3597, 3614, 3632, 3658, 3670, 3686, 3701, 3720, 3736, 3755, 3776, 3796, 3818, 3840, 3863, 3880, 3893, 3910, 3929, 3949, 3967, 3991, 4016, 4034, 4044, 4056, 4073, 4089, 4102, 4118, 4128, 4148, 4171, 4193, 4215, 4238, 4258, 4278, 4304, 4328, 4351, 4373, 4395, 4406, 4432, 4457, 4484, 4505, 4529, 4554, 4569, 4581, 4602, 4622, 4660, 4682, 4703, 4725, 4749, 4768, 4789, 4811, 4826, 4851, 4873, 4893, 4912, 4922, 4938, 4966, 4985, 5004, 5028, 5040, 5055, 5070, 5092, 5114, 5135, 5153, 5172, 5193, 5224, 5243, 5275, 5304, 5325, 5345, 5370, 5392, 5415, 5446, 5464, 5489, 5513, 5539, 5555, 5577, 5589}

func (i RuleName) String() string {
	if i >= RuleName(len(_RuleName_index)-1) {
		return "RuleName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RuleName_name[_RuleName_index[i]:_RuleName_index[i+1]]
}
