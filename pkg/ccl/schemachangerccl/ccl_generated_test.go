// Copyright 2022 The Cockroach Authors.
//
// Licensed as a CockroachDB Enterprise file under the Cockroach Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/cockroachdb/cockroach/blob/master/licenses/CCL.txt

// Code generated by sctestgen, DO NOT EDIT.

package schemachangerccl

import (
	"testing"

	"github.com/cockroachdb/cockroach/pkg/sql/schemachanger/sctest"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/cockroachdb/cockroach/pkg/util/log"
)

func TestBackupccl_create_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/ccl/schemachangerccl/testdata/end_to_end/create_index", newCluster)
}
func TestEndToEndSideEffectsccl_create_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.EndToEndSideEffects(t, "pkg/ccl/schemachangerccl/testdata/end_to_end/create_index", newCluster)
}
func TestGenerateSchemaChangeCorpusccl_create_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.GenerateSchemaChangeCorpus(t, "pkg/ccl/schemachangerccl/testdata/end_to_end/create_index", newCluster)
}
func TestPauseccl_create_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Pause(t, "pkg/ccl/schemachangerccl/testdata/end_to_end/create_index", newCluster)
}
func TestRollbackccl_create_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Rollback(t, "pkg/ccl/schemachangerccl/testdata/end_to_end/create_index", newCluster)
}
func TestBackupccl_drop_database_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_database_multiregion_primary_region", newCluster)
}
func TestEndToEndSideEffectsccl_drop_database_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.EndToEndSideEffects(t, "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_database_multiregion_primary_region", newCluster)
}
func TestGenerateSchemaChangeCorpusccl_drop_database_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.GenerateSchemaChangeCorpus(t, "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_database_multiregion_primary_region", newCluster)
}
func TestPauseccl_drop_database_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Pause(t, "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_database_multiregion_primary_region", newCluster)
}
func TestRollbackccl_drop_database_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Rollback(t, "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_database_multiregion_primary_region", newCluster)
}
func TestBackupccl_drop_table_multiregion(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion", newCluster)
}
func TestEndToEndSideEffectsccl_drop_table_multiregion(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.EndToEndSideEffects(t, "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion", newCluster)
}
func TestGenerateSchemaChangeCorpusccl_drop_table_multiregion(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.GenerateSchemaChangeCorpus(t, "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion", newCluster)
}
func TestPauseccl_drop_table_multiregion(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Pause(t, "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion", newCluster)
}
func TestRollbackccl_drop_table_multiregion(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Rollback(t, "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion", newCluster)
}
func TestBackupccl_drop_table_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion_primary_region", newCluster)
}
func TestEndToEndSideEffectsccl_drop_table_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.EndToEndSideEffects(t, "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion_primary_region", newCluster)
}
func TestGenerateSchemaChangeCorpusccl_drop_table_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.GenerateSchemaChangeCorpus(t, "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion_primary_region", newCluster)
}
func TestPauseccl_drop_table_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Pause(t, "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion_primary_region", newCluster)
}
func TestRollbackccl_drop_table_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Rollback(t, "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion_primary_region", newCluster)
}
