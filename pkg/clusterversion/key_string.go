// Code generated by "stringer"; DO NOT EDIT.

package clusterversion

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[V21_2-0]
	_ = x[Start22_1-1]
	_ = x[PebbleFormatBlockPropertyCollector-2]
	_ = x[ProbeRequest-3]
	_ = x[PublicSchemasWithDescriptors-4]
	_ = x[EnsureSpanConfigReconciliation-5]
	_ = x[EnsureSpanConfigSubscription-6]
	_ = x[EnableSpanConfigStore-7]
	_ = x[SCRAMAuthentication-8]
	_ = x[UnsafeLossOfQuorumRecoveryRangeLog-9]
	_ = x[AlterSystemProtectedTimestampAddColumn-10]
	_ = x[EnableProtectedTimestampsForTenant-11]
	_ = x[DeleteCommentsWithDroppedIndexes-12]
	_ = x[RemoveIncompatibleDatabasePrivileges-13]
	_ = x[AddRaftAppliedIndexTermMigration-14]
	_ = x[PostAddRaftAppliedIndexTermMigration-15]
	_ = x[DontProposeWriteTimestampForLeaseTransfers-16]
	_ = x[EnablePebbleFormatVersionBlockProperties-17]
	_ = x[MVCCIndexBackfiller-18]
	_ = x[EnableLeaseHolderRemoval-19]
	_ = x[LooselyCoupledRaftLogTruncation-20]
	_ = x[ChangefeedIdleness-21]
	_ = x[EnableDeclarativeSchemaChanger-22]
	_ = x[RowLevelTTL-23]
	_ = x[PebbleFormatSplitUserKeysMarked-24]
	_ = x[EnableNewStoreRebalancer-25]
	_ = x[ClusterLocksVirtualTable-26]
	_ = x[AutoStatsTableSettings-27]
	_ = x[SuperRegions-28]
	_ = x[EnableNewChangefeedOptions-29]
	_ = x[SpanCountTable-30]
	_ = x[PreSeedSpanCountTable-31]
	_ = x[SeedSpanCountTable-32]
	_ = x[V22_1-33]
	_ = x[Start22_2-34]
	_ = x[LocalTimestamps-35]
	_ = x[PebbleFormatSplitUserKeysMarkedCompacted-36]
	_ = x[EnsurePebbleFormatVersionRangeKeys-37]
	_ = x[EnablePebbleFormatVersionRangeKeys-38]
	_ = x[TrigramInvertedIndexes-39]
	_ = x[RemoveGrantPrivilege-40]
	_ = x[MVCCRangeTombstones-41]
	_ = x[UpgradeSequenceToBeReferencedByID-42]
	_ = x[SampledStmtDiagReqs-43]
	_ = x[AddSSTableTombstones-44]
	_ = x[SystemPrivilegesTable-45]
	_ = x[EnablePredicateProjectionChangefeed-46]
	_ = x[AlterSystemSQLInstancesAddLocality-47]
	_ = x[SystemExternalConnectionsTable-48]
	_ = x[AlterSystemStatementStatisticsAddIndexRecommendations-49]
	_ = x[RoleIDSequence-50]
	_ = x[AddSystemUserIDColumn-51]
	_ = x[UsersHaveIDs-52]
	_ = x[SetUserIDNotNull-53]
	_ = x[SQLSchemaTelemetryScheduledJobs-54]
}

const _Key_name = "V21_2Start22_1PebbleFormatBlockPropertyCollectorProbeRequestPublicSchemasWithDescriptorsEnsureSpanConfigReconciliationEnsureSpanConfigSubscriptionEnableSpanConfigStoreSCRAMAuthenticationUnsafeLossOfQuorumRecoveryRangeLogAlterSystemProtectedTimestampAddColumnEnableProtectedTimestampsForTenantDeleteCommentsWithDroppedIndexesRemoveIncompatibleDatabasePrivilegesAddRaftAppliedIndexTermMigrationPostAddRaftAppliedIndexTermMigrationDontProposeWriteTimestampForLeaseTransfersEnablePebbleFormatVersionBlockPropertiesMVCCIndexBackfillerEnableLeaseHolderRemovalLooselyCoupledRaftLogTruncationChangefeedIdlenessEnableDeclarativeSchemaChangerRowLevelTTLPebbleFormatSplitUserKeysMarkedEnableNewStoreRebalancerClusterLocksVirtualTableAutoStatsTableSettingsSuperRegionsEnableNewChangefeedOptionsSpanCountTablePreSeedSpanCountTableSeedSpanCountTableV22_1Start22_2LocalTimestampsPebbleFormatSplitUserKeysMarkedCompactedEnsurePebbleFormatVersionRangeKeysEnablePebbleFormatVersionRangeKeysTrigramInvertedIndexesRemoveGrantPrivilegeMVCCRangeTombstonesUpgradeSequenceToBeReferencedByIDSampledStmtDiagReqsAddSSTableTombstonesSystemPrivilegesTableEnablePredicateProjectionChangefeedAlterSystemSQLInstancesAddLocalitySystemExternalConnectionsTableAlterSystemStatementStatisticsAddIndexRecommendationsRoleIDSequenceAddSystemUserIDColumnUsersHaveIDsSetUserIDNotNullSQLSchemaTelemetryScheduledJobs"

var _Key_index = [...]uint16{0, 5, 14, 48, 60, 88, 118, 146, 167, 186, 220, 258, 292, 324, 360, 392, 428, 470, 510, 529, 553, 584, 602, 632, 643, 674, 698, 722, 744, 756, 782, 796, 817, 835, 840, 849, 864, 904, 938, 972, 994, 1014, 1033, 1066, 1085, 1105, 1126, 1161, 1195, 1225, 1278, 1292, 1313, 1325, 1341, 1372}

func (i Key) String() string {
	if i < 0 || i >= Key(len(_Key_index)-1) {
		return "Key(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Key_name[_Key_index[i]:_Key_index[i+1]]
}
