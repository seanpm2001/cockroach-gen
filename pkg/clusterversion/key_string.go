// Code generated by "stringer"; DO NOT EDIT.

package clusterversion

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[V21_2-0]
	_ = x[Start22_1-1]
	_ = x[ProbeRequest-2]
	_ = x[PublicSchemasWithDescriptors-3]
	_ = x[EnsureSpanConfigReconciliation-4]
	_ = x[EnsureSpanConfigSubscription-5]
	_ = x[EnableSpanConfigStore-6]
	_ = x[EnablePebbleFormatVersionBlockProperties-7]
	_ = x[ChangefeedIdleness-8]
	_ = x[EnableNewStoreRebalancer-9]
	_ = x[AutoStatsTableSettings-10]
	_ = x[EnableNewChangefeedOptions-11]
	_ = x[V22_1-12]
	_ = x[Start22_2-13]
	_ = x[LocalTimestamps-14]
	_ = x[PebbleFormatSplitUserKeysMarkedCompacted-15]
	_ = x[EnsurePebbleFormatVersionRangeKeys-16]
	_ = x[EnablePebbleFormatVersionRangeKeys-17]
	_ = x[TrigramInvertedIndexes-18]
	_ = x[RemoveGrantPrivilege-19]
	_ = x[MVCCRangeTombstones-20]
	_ = x[UpgradeSequenceToBeReferencedByID-21]
	_ = x[SampledStmtDiagReqs-22]
	_ = x[AddSSTableTombstones-23]
	_ = x[SystemPrivilegesTable-24]
	_ = x[EnablePredicateProjectionChangefeed-25]
	_ = x[AlterSystemSQLInstancesAddLocality-26]
	_ = x[SystemExternalConnectionsTable-27]
	_ = x[AlterSystemStatementStatisticsAddIndexRecommendations-28]
	_ = x[RoleIDSequence-29]
	_ = x[AddSystemUserIDColumn-30]
	_ = x[SystemUsersIDColumnIsBackfilled-31]
	_ = x[SetSystemUsersUserIDColumnNotNull-32]
	_ = x[SQLSchemaTelemetryScheduledJobs-33]
	_ = x[SchemaChangeSupportsCreateFunction-34]
	_ = x[DeleteRequestReturnKey-35]
	_ = x[PebbleFormatPrePebblev1Marked-36]
	_ = x[RoleOptionsTableHasIDColumn-37]
	_ = x[RoleOptionsIDColumnIsBackfilled-38]
	_ = x[SetRoleOptionsUserIDColumnNotNull-39]
	_ = x[UseDelRangeInGCJob-40]
	_ = x[WaitedForDelRangeInGCJob-41]
	_ = x[RangefeedUseOneStreamPerNode-42]
}

const _Key_name = "V21_2Start22_1ProbeRequestPublicSchemasWithDescriptorsEnsureSpanConfigReconciliationEnsureSpanConfigSubscriptionEnableSpanConfigStoreEnablePebbleFormatVersionBlockPropertiesChangefeedIdlenessEnableNewStoreRebalancerAutoStatsTableSettingsEnableNewChangefeedOptionsV22_1Start22_2LocalTimestampsPebbleFormatSplitUserKeysMarkedCompactedEnsurePebbleFormatVersionRangeKeysEnablePebbleFormatVersionRangeKeysTrigramInvertedIndexesRemoveGrantPrivilegeMVCCRangeTombstonesUpgradeSequenceToBeReferencedByIDSampledStmtDiagReqsAddSSTableTombstonesSystemPrivilegesTableEnablePredicateProjectionChangefeedAlterSystemSQLInstancesAddLocalitySystemExternalConnectionsTableAlterSystemStatementStatisticsAddIndexRecommendationsRoleIDSequenceAddSystemUserIDColumnSystemUsersIDColumnIsBackfilledSetSystemUsersUserIDColumnNotNullSQLSchemaTelemetryScheduledJobsSchemaChangeSupportsCreateFunctionDeleteRequestReturnKeyPebbleFormatPrePebblev1MarkedRoleOptionsTableHasIDColumnRoleOptionsIDColumnIsBackfilledSetRoleOptionsUserIDColumnNotNullUseDelRangeInGCJobWaitedForDelRangeInGCJobRangefeedUseOneStreamPerNode"

var _Key_index = [...]uint16{0, 5, 14, 26, 54, 84, 112, 133, 173, 191, 215, 237, 263, 268, 277, 292, 332, 366, 400, 422, 442, 461, 494, 513, 533, 554, 589, 623, 653, 706, 720, 741, 772, 805, 836, 870, 892, 921, 948, 979, 1012, 1030, 1054, 1082}

func (i Key) String() string {
	if i < 0 || i >= Key(len(_Key_index)-1) {
		return "Key(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Key_name[_Key_index[i]:_Key_index[i+1]]
}
