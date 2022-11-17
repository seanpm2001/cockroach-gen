// Code generated by "stringer"; DO NOT EDIT.

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
	_ = x[NormCycleTestRelTrueToFalse-53]
	_ = x[NormCycleTestRelFalseToTrue-54]
	_ = x[DecorrelateJoin-55]
	_ = x[DecorrelateProjectSet-56]
	_ = x[TryDecorrelateSelect-57]
	_ = x[TryDecorrelateProject-58]
	_ = x[TryDecorrelateProjectSelect-59]
	_ = x[TryDecorrelateProjectInnerJoin-60]
	_ = x[TryDecorrelateInnerJoin-61]
	_ = x[TryDecorrelateInnerLeftJoin-62]
	_ = x[TryDecorrelateGroupBy-63]
	_ = x[TryDecorrelateScalarGroupBy-64]
	_ = x[TryDecorrelateSemiJoin-65]
	_ = x[TryDecorrelateLimitOne-66]
	_ = x[TryDecorrelateLimit-67]
	_ = x[TryDecorrelateProjectSet-68]
	_ = x[TryDecorrelateWindow-69]
	_ = x[TryDecorrelateMax1Row-70]
	_ = x[HoistSelectExists-71]
	_ = x[HoistSelectNotExists-72]
	_ = x[HoistSelectSubquery-73]
	_ = x[HoistProjectSubquery-74]
	_ = x[HoistJoinSubquery-75]
	_ = x[HoistValuesSubquery-76]
	_ = x[HoistProjectSetSubquery-77]
	_ = x[NormalizeSelectAnyFilter-78]
	_ = x[NormalizeJoinAnyFilter-79]
	_ = x[NormalizeSelectNotAnyFilter-80]
	_ = x[NormalizeJoinNotAnyFilter-81]
	_ = x[FoldNullCast-82]
	_ = x[FoldNullUnary-83]
	_ = x[FoldNullBinaryLeft-84]
	_ = x[FoldNullBinaryRight-85]
	_ = x[FoldNullInNonEmpty-86]
	_ = x[FoldInEmpty-87]
	_ = x[FoldNotInEmpty-88]
	_ = x[FoldArray-89]
	_ = x[FoldBinary-90]
	_ = x[FoldUnary-91]
	_ = x[FoldComparison-92]
	_ = x[FoldCast-93]
	_ = x[FoldAssignmentCast-94]
	_ = x[FoldIndirection-95]
	_ = x[FoldColumnAccess-96]
	_ = x[FoldFunctionWithNullArg-97]
	_ = x[FoldFunction-98]
	_ = x[FoldEqualsAnyNull-99]
	_ = x[ConvertGroupByToDistinct-100]
	_ = x[EliminateGroupByProject-101]
	_ = x[EliminateJoinUnderGroupByLeft-102]
	_ = x[EliminateJoinUnderGroupByRight-103]
	_ = x[EliminateDistinct-104]
	_ = x[ReduceGroupingCols-105]
	_ = x[ReduceNotNullGroupingCols-106]
	_ = x[EliminateAggDistinctForKeys-107]
	_ = x[EliminateAggFilteredDistinctForKeys-108]
	_ = x[EliminateDistinctNoColumns-109]
	_ = x[EliminateEnsureDistinctNoColumns-110]
	_ = x[EliminateDistinctOnValues-111]
	_ = x[PushAggDistinctIntoGroupBy-112]
	_ = x[PushAggFilterIntoScalarGroupBy-113]
	_ = x[ConvertCountToCountRows-114]
	_ = x[ConvertRegressionCountToCount-115]
	_ = x[FoldGroupingOperators-116]
	_ = x[InlineConstVar-117]
	_ = x[InlineProjectConstants-118]
	_ = x[InlineSelectConstants-119]
	_ = x[InlineJoinConstantsLeft-120]
	_ = x[InlineJoinConstantsRight-121]
	_ = x[PushSelectIntoInlinableProject-122]
	_ = x[InlineSelectVirtualColumns-123]
	_ = x[InlineProjectInProject-124]
	_ = x[CommuteRightJoin-125]
	_ = x[SimplifyJoinFilters-126]
	_ = x[DetectJoinContradiction-127]
	_ = x[PushFilterIntoJoinLeftAndRight-128]
	_ = x[MapFilterIntoJoinLeft-129]
	_ = x[MapFilterIntoJoinRight-130]
	_ = x[MapEqualityIntoJoinLeftAndRight-131]
	_ = x[PushFilterIntoJoinLeft-132]
	_ = x[PushFilterIntoJoinRight-133]
	_ = x[SimplifyLeftJoin-134]
	_ = x[SimplifyRightJoin-135]
	_ = x[EliminateSemiJoin-136]
	_ = x[SimplifyZeroCardinalitySemiJoin-137]
	_ = x[EliminateAntiJoin-138]
	_ = x[SimplifyZeroCardinalityAntiJoin-139]
	_ = x[EliminateJoinNoColsLeft-140]
	_ = x[EliminateJoinNoColsRight-141]
	_ = x[HoistJoinProjectRight-142]
	_ = x[HoistJoinProjectLeft-143]
	_ = x[SimplifyJoinNotNullEquality-144]
	_ = x[ExtractJoinComparisons-145]
	_ = x[SortFiltersInJoin-146]
	_ = x[LeftAssociateJoinsLeft-147]
	_ = x[LeftAssociateJoinsRight-148]
	_ = x[RightAssociateJoinsLeft-149]
	_ = x[RightAssociateJoinsRight-150]
	_ = x[RemoveJoinNotNullCondition-151]
	_ = x[ProjectInnerJoinValues-152]
	_ = x[EliminateLimit-153]
	_ = x[EliminateOffset-154]
	_ = x[PushLimitIntoProject-155]
	_ = x[PushOffsetIntoProject-156]
	_ = x[PushLimitIntoOffset-157]
	_ = x[PushLimitIntoOrdinality-158]
	_ = x[PushLimitIntoJoinLeft-159]
	_ = x[PushLimitIntoJoinRight-160]
	_ = x[FoldLimits-161]
	_ = x[AssociateLimitJoinsLeft-162]
	_ = x[AssociateLimitJoinsRight-163]
	_ = x[EliminateMax1Row-164]
	_ = x[SimplifyPartialIndexProjections-165]
	_ = x[FoldPlusZero-166]
	_ = x[FoldZeroPlus-167]
	_ = x[FoldMinusZero-168]
	_ = x[FoldMultOne-169]
	_ = x[FoldOneMult-170]
	_ = x[FoldDivOne-171]
	_ = x[FoldFloorDivOne-172]
	_ = x[InvertMinus-173]
	_ = x[EliminateUnaryMinus-174]
	_ = x[SimplifyLimitOrdering-175]
	_ = x[SimplifyOffsetOrdering-176]
	_ = x[SimplifyGroupByOrdering-177]
	_ = x[SimplifyOrdinalityOrdering-178]
	_ = x[SimplifyExplainOrdering-179]
	_ = x[SimplifyWithBindingOrdering-180]
	_ = x[EliminateJoinUnderProjectLeft-181]
	_ = x[EliminateJoinUnderProjectRight-182]
	_ = x[EliminateProject-183]
	_ = x[MergeProjects-184]
	_ = x[MergeProjectWithValues-185]
	_ = x[PushColumnRemappingIntoValues-186]
	_ = x[PushAssignmentCastsIntoValues-187]
	_ = x[FoldTupleAccessIntoValues-188]
	_ = x[FoldJSONAccessIntoValues-189]
	_ = x[ConvertZipArraysToValues-190]
	_ = x[PruneProjectCols-191]
	_ = x[PruneScanCols-192]
	_ = x[PruneSelectCols-193]
	_ = x[PruneLimitCols-194]
	_ = x[PruneOffsetCols-195]
	_ = x[PruneJoinLeftCols-196]
	_ = x[PruneJoinRightCols-197]
	_ = x[PruneSemiAntiJoinRightCols-198]
	_ = x[PruneAggCols-199]
	_ = x[PruneGroupByCols-200]
	_ = x[PruneValuesCols-201]
	_ = x[PruneOrdinalityCols-202]
	_ = x[PruneExplainCols-203]
	_ = x[PruneProjectSetCols-204]
	_ = x[PruneWindowOutputCols-205]
	_ = x[PruneWindowInputCols-206]
	_ = x[PruneMutationFetchCols-207]
	_ = x[PruneMutationInputCols-208]
	_ = x[PruneMutationReturnCols-209]
	_ = x[PruneWithScanCols-210]
	_ = x[PruneWithCols-211]
	_ = x[PruneUnionAllCols-212]
	_ = x[RejectNullsLeftJoin-213]
	_ = x[RejectNullsRightJoin-214]
	_ = x[RejectNullsGroupBy-215]
	_ = x[RejectNullsUnderJoinLeft-216]
	_ = x[RejectNullsUnderJoinRight-217]
	_ = x[RejectNullsProject-218]
	_ = x[CommuteVar-219]
	_ = x[CommuteConst-220]
	_ = x[EliminateCoalesce-221]
	_ = x[SimplifyCoalesce-222]
	_ = x[EliminateCast-223]
	_ = x[InlineAnyWithScan-224]
	_ = x[NormalizeInConst-225]
	_ = x[FoldInNull-226]
	_ = x[SimplifyInSingleElement-227]
	_ = x[SimplifyNotInSingleElement-228]
	_ = x[UnifyComparisonTypes-229]
	_ = x[EliminateExistsZeroRows-230]
	_ = x[EliminateExistsProject-231]
	_ = x[EliminateExistsGroupBy-232]
	_ = x[InlineExistsSelectTuple-233]
	_ = x[IntroduceExistsLimit-234]
	_ = x[EliminateExistsLimit-235]
	_ = x[SimplifyCaseWhenConstValue-236]
	_ = x[InlineAnyValuesSingleCol-237]
	_ = x[InlineAnyValuesMultiCol-238]
	_ = x[SimplifyEqualsAnyTuple-239]
	_ = x[SimplifyAnyScalarArray-240]
	_ = x[FoldCollate-241]
	_ = x[NormalizeArrayFlattenToAgg-242]
	_ = x[SimplifySameVarEqualities-243]
	_ = x[SimplifySameVarInequalities-244]
	_ = x[SimplifyNotDisjoint-245]
	_ = x[ConvertJSONSubscriptToFetchValue-246]
	_ = x[SimplifySelectFilters-247]
	_ = x[ConsolidateSelectFilters-248]
	_ = x[DeduplicateSelectFilters-249]
	_ = x[EliminateSelect-250]
	_ = x[MergeSelects-251]
	_ = x[PushSelectIntoProject-252]
	_ = x[MergeSelectInnerJoin-253]
	_ = x[PushSelectCondLeftIntoJoinLeftAndRight-254]
	_ = x[PushSelectIntoJoinLeft-255]
	_ = x[PushSelectIntoGroupBy-256]
	_ = x[RemoveNotNullCondition-257]
	_ = x[SimplifyIsNullCondition-258]
	_ = x[PushSelectIntoProjectSet-259]
	_ = x[PushFilterIntoSetOp-260]
	_ = x[EliminateSetLeft-261]
	_ = x[EliminateSetRight-262]
	_ = x[EliminateDistinctSetLeft-263]
	_ = x[EliminateDistinctSetRight-264]
	_ = x[SimplifyExcept-265]
	_ = x[SimplifyIntersectLeft-266]
	_ = x[SimplifyIntersectRight-267]
	_ = x[ConvertUnionToDistinctUnionAll-268]
	_ = x[EliminateWindow-269]
	_ = x[ReduceWindowPartitionCols-270]
	_ = x[SimplifyWindowOrdering-271]
	_ = x[PushSelectIntoWindow-272]
	_ = x[PushLimitIntoWindow-273]
	_ = x[InlineWith-274]
	_ = x[ApplyLimitToRecursiveCTEScan-275]
	_ = x[TryAddLimitToRecursiveBranch-276]
	_ = x[startExploreRule-277]
	_ = x[MemoCycleTestRelRule-278]
	_ = x[MemoCycleTestRelRuleFilter-279]
	_ = x[ReplaceScalarMinMaxWithScalarSubqueries-280]
	_ = x[ReplaceFilteredScalarMinMaxWithSubqueries-281]
	_ = x[ReplaceScalarMinMaxWithLimit-282]
	_ = x[ReplaceMinWithLimit-283]
	_ = x[ReplaceMaxWithLimit-284]
	_ = x[GenerateStreamingGroupBy-285]
	_ = x[SplitGroupByScanIntoUnionScans-286]
	_ = x[SplitGroupByFilteredScanIntoUnionScans-287]
	_ = x[EliminateIndexJoinOrProjectInsideGroupBy-288]
	_ = x[GenerateLimitedGroupByScans-289]
	_ = x[ReorderJoins-290]
	_ = x[CommuteLeftJoin-291]
	_ = x[CommuteSemiJoin-292]
	_ = x[ConvertSemiToInnerJoin-293]
	_ = x[SplitDisjunctionOfJoinTerms-294]
	_ = x[SplitDisjunctionOfAntiJoinTerms-295]
	_ = x[GenerateMergeJoins-296]
	_ = x[GenerateLookupJoins-297]
	_ = x[GenerateInvertedJoins-298]
	_ = x[GenerateInvertedJoinsFromSelect-299]
	_ = x[GenerateLookupJoinsWithFilter-300]
	_ = x[GenerateLookupJoinsWithVirtualCols-301]
	_ = x[GenerateLookupJoinsWithVirtualColsAndFilter-302]
	_ = x[PushJoinIntoIndexJoin-303]
	_ = x[HoistProjectFromInnerJoin-304]
	_ = x[HoistProjectFromLeftJoin-305]
	_ = x[GenerateLocalityOptimizedAntiJoin-306]
	_ = x[GenerateLocalityOptimizedLookupJoin-307]
	_ = x[GenerateLimitedScans-308]
	_ = x[PushLimitIntoFilteredScan-309]
	_ = x[PushLimitIntoIndexJoin-310]
	_ = x[SplitLimitedScanIntoUnionScans-311]
	_ = x[SplitLimitedSelectIntoUnionSelects-312]
	_ = x[GenerateTopK-313]
	_ = x[GenerateLimitedTopKScans-314]
	_ = x[GeneratePartialOrderTopK-315]
	_ = x[EliminateIndexJoinInsideProject-316]
	_ = x[GenerateIndexScans-317]
	_ = x[GenerateLocalityOptimizedScan-318]
	_ = x[GeneratePartialIndexScans-319]
	_ = x[GenerateConstrainedScans-320]
	_ = x[GenerateInvertedIndexScans-321]
	_ = x[GenerateZigzagJoins-322]
	_ = x[GenerateInvertedIndexZigzagJoins-323]
	_ = x[SplitDisjunction-324]
	_ = x[SplitDisjunctionAddKey-325]
	_ = x[GenerateStreamingSetOp-326]
	_ = x[NumRuleNames-327]
}

const _RuleName_name = "InvalidRuleNameSimplifyRootOrderingPruneRootColsSimplifyZeroCardinalityGroupNumManualRuleNamesEliminateAggDistinctNormalizeNestedAndsSimplifyTrueAndSimplifyAndTrueSimplifyFalseAndSimplifyAndFalseSimplifyTrueOrSimplifyOrTrueSimplifyFalseOrSimplifyOrFalseSimplifyRangeFoldNullAndOrFoldNotTrueFoldNotFalseFoldNotNullNegateComparisonEliminateNotNegateAndNegateOrExtractRedundantConjunctCommuteVarInequalityCommuteConstInequalityNormalizeCmpPlusConstNormalizeCmpMinusConstNormalizeCmpConstMinusNormalizeTupleEqualityFoldNullComparisonLeftFoldNullComparisonRightFoldIsNullFoldNonNullIsNullFoldNullTupleIsTupleNullFoldNonNullTupleIsTupleNullFoldIsNotNullFoldNonNullIsNotNullFoldNonNullTupleIsTupleNotNullFoldNullTupleIsTupleNotNullCommuteNullIsNormalizeCmpTimeZoneFunctionNormalizeCmpTimeZoneFunctionTZFoldEqZeroSTDistanceFoldCmpSTDistanceLeftFoldCmpSTDistanceRightFoldCmpSTMaxDistanceLeftFoldCmpSTMaxDistanceRightFoldEqTrueFoldEqFalseFoldNeTrueFoldNeFalseNormCycleTestRelTrueToFalseNormCycleTestRelFalseToTrueDecorrelateJoinDecorrelateProjectSetTryDecorrelateSelectTryDecorrelateProjectTryDecorrelateProjectSelectTryDecorrelateProjectInnerJoinTryDecorrelateInnerJoinTryDecorrelateInnerLeftJoinTryDecorrelateGroupByTryDecorrelateScalarGroupByTryDecorrelateSemiJoinTryDecorrelateLimitOneTryDecorrelateLimitTryDecorrelateProjectSetTryDecorrelateWindowTryDecorrelateMax1RowHoistSelectExistsHoistSelectNotExistsHoistSelectSubqueryHoistProjectSubqueryHoistJoinSubqueryHoistValuesSubqueryHoistProjectSetSubqueryNormalizeSelectAnyFilterNormalizeJoinAnyFilterNormalizeSelectNotAnyFilterNormalizeJoinNotAnyFilterFoldNullCastFoldNullUnaryFoldNullBinaryLeftFoldNullBinaryRightFoldNullInNonEmptyFoldInEmptyFoldNotInEmptyFoldArrayFoldBinaryFoldUnaryFoldComparisonFoldCastFoldAssignmentCastFoldIndirectionFoldColumnAccessFoldFunctionWithNullArgFoldFunctionFoldEqualsAnyNullConvertGroupByToDistinctEliminateGroupByProjectEliminateJoinUnderGroupByLeftEliminateJoinUnderGroupByRightEliminateDistinctReduceGroupingColsReduceNotNullGroupingColsEliminateAggDistinctForKeysEliminateAggFilteredDistinctForKeysEliminateDistinctNoColumnsEliminateEnsureDistinctNoColumnsEliminateDistinctOnValuesPushAggDistinctIntoGroupByPushAggFilterIntoScalarGroupByConvertCountToCountRowsConvertRegressionCountToCountFoldGroupingOperatorsInlineConstVarInlineProjectConstantsInlineSelectConstantsInlineJoinConstantsLeftInlineJoinConstantsRightPushSelectIntoInlinableProjectInlineSelectVirtualColumnsInlineProjectInProjectCommuteRightJoinSimplifyJoinFiltersDetectJoinContradictionPushFilterIntoJoinLeftAndRightMapFilterIntoJoinLeftMapFilterIntoJoinRightMapEqualityIntoJoinLeftAndRightPushFilterIntoJoinLeftPushFilterIntoJoinRightSimplifyLeftJoinSimplifyRightJoinEliminateSemiJoinSimplifyZeroCardinalitySemiJoinEliminateAntiJoinSimplifyZeroCardinalityAntiJoinEliminateJoinNoColsLeftEliminateJoinNoColsRightHoistJoinProjectRightHoistJoinProjectLeftSimplifyJoinNotNullEqualityExtractJoinComparisonsSortFiltersInJoinLeftAssociateJoinsLeftLeftAssociateJoinsRightRightAssociateJoinsLeftRightAssociateJoinsRightRemoveJoinNotNullConditionProjectInnerJoinValuesEliminateLimitEliminateOffsetPushLimitIntoProjectPushOffsetIntoProjectPushLimitIntoOffsetPushLimitIntoOrdinalityPushLimitIntoJoinLeftPushLimitIntoJoinRightFoldLimitsAssociateLimitJoinsLeftAssociateLimitJoinsRightEliminateMax1RowSimplifyPartialIndexProjectionsFoldPlusZeroFoldZeroPlusFoldMinusZeroFoldMultOneFoldOneMultFoldDivOneFoldFloorDivOneInvertMinusEliminateUnaryMinusSimplifyLimitOrderingSimplifyOffsetOrderingSimplifyGroupByOrderingSimplifyOrdinalityOrderingSimplifyExplainOrderingSimplifyWithBindingOrderingEliminateJoinUnderProjectLeftEliminateJoinUnderProjectRightEliminateProjectMergeProjectsMergeProjectWithValuesPushColumnRemappingIntoValuesPushAssignmentCastsIntoValuesFoldTupleAccessIntoValuesFoldJSONAccessIntoValuesConvertZipArraysToValuesPruneProjectColsPruneScanColsPruneSelectColsPruneLimitColsPruneOffsetColsPruneJoinLeftColsPruneJoinRightColsPruneSemiAntiJoinRightColsPruneAggColsPruneGroupByColsPruneValuesColsPruneOrdinalityColsPruneExplainColsPruneProjectSetColsPruneWindowOutputColsPruneWindowInputColsPruneMutationFetchColsPruneMutationInputColsPruneMutationReturnColsPruneWithScanColsPruneWithColsPruneUnionAllColsRejectNullsLeftJoinRejectNullsRightJoinRejectNullsGroupByRejectNullsUnderJoinLeftRejectNullsUnderJoinRightRejectNullsProjectCommuteVarCommuteConstEliminateCoalesceSimplifyCoalesceEliminateCastInlineAnyWithScanNormalizeInConstFoldInNullSimplifyInSingleElementSimplifyNotInSingleElementUnifyComparisonTypesEliminateExistsZeroRowsEliminateExistsProjectEliminateExistsGroupByInlineExistsSelectTupleIntroduceExistsLimitEliminateExistsLimitSimplifyCaseWhenConstValueInlineAnyValuesSingleColInlineAnyValuesMultiColSimplifyEqualsAnyTupleSimplifyAnyScalarArrayFoldCollateNormalizeArrayFlattenToAggSimplifySameVarEqualitiesSimplifySameVarInequalitiesSimplifyNotDisjointConvertJSONSubscriptToFetchValueSimplifySelectFiltersConsolidateSelectFiltersDeduplicateSelectFiltersEliminateSelectMergeSelectsPushSelectIntoProjectMergeSelectInnerJoinPushSelectCondLeftIntoJoinLeftAndRightPushSelectIntoJoinLeftPushSelectIntoGroupByRemoveNotNullConditionSimplifyIsNullConditionPushSelectIntoProjectSetPushFilterIntoSetOpEliminateSetLeftEliminateSetRightEliminateDistinctSetLeftEliminateDistinctSetRightSimplifyExceptSimplifyIntersectLeftSimplifyIntersectRightConvertUnionToDistinctUnionAllEliminateWindowReduceWindowPartitionColsSimplifyWindowOrderingPushSelectIntoWindowPushLimitIntoWindowInlineWithApplyLimitToRecursiveCTEScanTryAddLimitToRecursiveBranchstartExploreRuleMemoCycleTestRelRuleMemoCycleTestRelRuleFilterReplaceScalarMinMaxWithScalarSubqueriesReplaceFilteredScalarMinMaxWithSubqueriesReplaceScalarMinMaxWithLimitReplaceMinWithLimitReplaceMaxWithLimitGenerateStreamingGroupBySplitGroupByScanIntoUnionScansSplitGroupByFilteredScanIntoUnionScansEliminateIndexJoinOrProjectInsideGroupByGenerateLimitedGroupByScansReorderJoinsCommuteLeftJoinCommuteSemiJoinConvertSemiToInnerJoinSplitDisjunctionOfJoinTermsSplitDisjunctionOfAntiJoinTermsGenerateMergeJoinsGenerateLookupJoinsGenerateInvertedJoinsGenerateInvertedJoinsFromSelectGenerateLookupJoinsWithFilterGenerateLookupJoinsWithVirtualColsGenerateLookupJoinsWithVirtualColsAndFilterPushJoinIntoIndexJoinHoistProjectFromInnerJoinHoistProjectFromLeftJoinGenerateLocalityOptimizedAntiJoinGenerateLocalityOptimizedLookupJoinGenerateLimitedScansPushLimitIntoFilteredScanPushLimitIntoIndexJoinSplitLimitedScanIntoUnionScansSplitLimitedSelectIntoUnionSelectsGenerateTopKGenerateLimitedTopKScansGeneratePartialOrderTopKEliminateIndexJoinInsideProjectGenerateIndexScansGenerateLocalityOptimizedScanGeneratePartialIndexScansGenerateConstrainedScansGenerateInvertedIndexScansGenerateZigzagJoinsGenerateInvertedIndexZigzagJoinsSplitDisjunctionSplitDisjunctionAddKeyGenerateStreamingSetOpNumRuleNames"

var _RuleName_index = [...]uint16{0, 15, 35, 48, 76, 94, 114, 133, 148, 163, 179, 195, 209, 223, 238, 253, 266, 279, 290, 302, 313, 329, 341, 350, 358, 382, 402, 424, 445, 467, 489, 511, 533, 556, 566, 583, 607, 634, 647, 667, 697, 724, 737, 765, 795, 815, 836, 858, 882, 907, 917, 928, 938, 949, 976, 1003, 1018, 1039, 1059, 1080, 1107, 1137, 1160, 1187, 1208, 1235, 1257, 1279, 1298, 1322, 1342, 1363, 1380, 1400, 1419, 1439, 1456, 1475, 1498, 1522, 1544, 1571, 1596, 1608, 1621, 1639, 1658, 1676, 1687, 1701, 1710, 1720, 1729, 1743, 1751, 1769, 1784, 1800, 1823, 1835, 1852, 1876, 1899, 1928, 1958, 1975, 1993, 2018, 2045, 2080, 2106, 2138, 2163, 2189, 2219, 2242, 2271, 2292, 2306, 2328, 2349, 2372, 2396, 2426, 2452, 2474, 2490, 2509, 2532, 2562, 2583, 2605, 2636, 2658, 2681, 2697, 2714, 2731, 2762, 2779, 2810, 2833, 2857, 2878, 2898, 2925, 2947, 2964, 2986, 3009, 3032, 3056, 3082, 3104, 3118, 3133, 3153, 3174, 3193, 3216, 3237, 3259, 3269, 3292, 3316, 3332, 3363, 3375, 3387, 3400, 3411, 3422, 3432, 3447, 3458, 3477, 3498, 3520, 3543, 3569, 3592, 3619, 3648, 3678, 3694, 3707, 3729, 3758, 3787, 3812, 3836, 3860, 3876, 3889, 3904, 3918, 3933, 3950, 3968, 3994, 4006, 4022, 4037, 4056, 4072, 4091, 4112, 4132, 4154, 4176, 4199, 4216, 4229, 4246, 4265, 4285, 4303, 4327, 4352, 4370, 4380, 4392, 4409, 4425, 4438, 4455, 4471, 4481, 4504, 4530, 4550, 4573, 4595, 4617, 4640, 4660, 4680, 4706, 4730, 4753, 4775, 4797, 4808, 4834, 4859, 4886, 4905, 4937, 4958, 4982, 5006, 5021, 5033, 5054, 5074, 5112, 5134, 5155, 5177, 5200, 5224, 5243, 5259, 5276, 5300, 5325, 5339, 5360, 5382, 5412, 5427, 5452, 5474, 5494, 5513, 5523, 5551, 5579, 5595, 5615, 5641, 5680, 5721, 5749, 5768, 5787, 5811, 5841, 5879, 5919, 5946, 5958, 5973, 5988, 6010, 6037, 6068, 6086, 6105, 6126, 6157, 6186, 6220, 6263, 6284, 6309, 6333, 6366, 6401, 6421, 6446, 6468, 6498, 6532, 6544, 6568, 6592, 6623, 6641, 6670, 6695, 6719, 6745, 6764, 6796, 6812, 6834, 6856, 6868}

func (i RuleName) String() string {
	if i >= RuleName(len(_RuleName_index)-1) {
		return "RuleName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RuleName_name[_RuleName_index[i]:_RuleName_index[i+1]]
}
