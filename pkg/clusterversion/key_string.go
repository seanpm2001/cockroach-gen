// Code generated by "stringer -type=Key"; DO NOT EDIT.

package clusterversion

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[replacedTruncatedAndRangeAppliedStateMigration-0]
	_ = x[replacedPostTruncatedAndRangeAppliedStateMigration-1]
	_ = x[TruncatedAndRangeAppliedStateMigration-2]
	_ = x[PostTruncatedAndRangeAppliedStateMigration-3]
	_ = x[V21_1-4]
	_ = x[Start21_1PLUS-5]
	_ = x[Start21_2-6]
	_ = x[JoinTokensTable-7]
	_ = x[AcquisitionTypeInLeaseHistory-8]
	_ = x[SerializeViewUDTs-9]
	_ = x[ExpressionIndexes-10]
	_ = x[DeleteDeprecatedNamespaceTableDescriptorMigration-11]
	_ = x[FixDescriptors-12]
	_ = x[DatabaseRoleSettings-13]
	_ = x[TenantUsageTable-14]
	_ = x[SQLInstancesTable-15]
	_ = x[NewRetryableRangefeedErrors-16]
	_ = x[AlterSystemWebSessionsCreateIndexes-17]
	_ = x[SeparatedIntentsMigration-18]
	_ = x[PostSeparatedIntentsMigration-19]
	_ = x[RetryJobsWithExponentialBackoff-20]
	_ = x[RecordsBasedRegistry-21]
	_ = x[AutoSpanConfigReconciliationJob-22]
	_ = x[PreventNewInterleavedTables-23]
	_ = x[EnsureNoInterleavedTables-24]
	_ = x[DefaultPrivileges-25]
	_ = x[ZonesTableForSecondaryTenants-26]
	_ = x[UseKeyEncodeForHashShardedIndexes-27]
	_ = x[DatabasePlacementPolicy-28]
	_ = x[GeneratedAsIdentity-29]
	_ = x[OnUpdateExpressions-30]
	_ = x[SpanConfigurationsTable-31]
	_ = x[BoundedStaleness-32]
	_ = x[DateAndIntervalStyle-33]
	_ = x[PebbleFormatVersioned-34]
	_ = x[MarkerDataKeysRegistry-35]
	_ = x[PebbleSetWithDelete-36]
	_ = x[TenantUsageSingleConsumptionColumn-37]
	_ = x[SQLStatsTables-38]
	_ = x[SQLStatsCompactionScheduledJob-39]
	_ = x[V21_2-40]
}

const _Key_name = "replacedTruncatedAndRangeAppliedStateMigrationreplacedPostTruncatedAndRangeAppliedStateMigrationTruncatedAndRangeAppliedStateMigrationPostTruncatedAndRangeAppliedStateMigrationV21_1Start21_1PLUSStart21_2JoinTokensTableAcquisitionTypeInLeaseHistorySerializeViewUDTsExpressionIndexesDeleteDeprecatedNamespaceTableDescriptorMigrationFixDescriptorsDatabaseRoleSettingsTenantUsageTableSQLInstancesTableNewRetryableRangefeedErrorsAlterSystemWebSessionsCreateIndexesSeparatedIntentsMigrationPostSeparatedIntentsMigrationRetryJobsWithExponentialBackoffRecordsBasedRegistryAutoSpanConfigReconciliationJobPreventNewInterleavedTablesEnsureNoInterleavedTablesDefaultPrivilegesZonesTableForSecondaryTenantsUseKeyEncodeForHashShardedIndexesDatabasePlacementPolicyGeneratedAsIdentityOnUpdateExpressionsSpanConfigurationsTableBoundedStalenessDateAndIntervalStylePebbleFormatVersionedMarkerDataKeysRegistryPebbleSetWithDeleteTenantUsageSingleConsumptionColumnSQLStatsTablesSQLStatsCompactionScheduledJobV21_2"

var _Key_index = [...]uint16{0, 46, 96, 134, 176, 181, 194, 203, 218, 247, 264, 281, 330, 344, 364, 380, 397, 424, 459, 484, 513, 544, 564, 595, 622, 647, 664, 693, 726, 749, 768, 787, 810, 826, 846, 867, 889, 908, 942, 956, 986, 991}

func (i Key) String() string {
	if i < 0 || i >= Key(len(_Key_index)-1) {
		return "Key(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Key_name[_Key_index[i]:_Key_index[i+1]]
}
