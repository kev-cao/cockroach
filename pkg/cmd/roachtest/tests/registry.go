// Copyright 2018 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

package tests

import (
	"github.com/cockroachdb/cockroach/pkg/cmd/roachtest/registry"
	"github.com/cockroachdb/cockroach/pkg/cmd/roachtest/tests/perturbation"
)

// RegisterTests registers all tests to the Registry. This powers `roachtest run`.
func RegisterTests(r registry.Registry) {
	registerAWSDMS(r)
	registerAcceptance(r)
	registerActiveRecord(r)
	registerAdmission(r)
	registerAllocator(r)
	registerLimitCapacity(r)
	registerAllocationBench(r)
	registerAlterPK(r)
	registerAsyncpg(r)
	registerBackup(r)
	registerBackupMixedVersion(r)
	registerBackupNodeShutdown(r)
	registerBackupRestoreRoundTrip(r)
	registerBackupFixtures(r)
	registerBackupS3Clones(r)
	registerCDC(r)
	registerCDCBench(r)
	registerCDCFiltering(r)
	registerCDCMixedVersions(r)
	registerExportParquet(r)
	registerCancel(r)
	registerChangeReplicasMixedVersion(r)
	registerClearRange(r)
	registerClockJumpTests(r)
	registerClockMonotonicTests(r)
	registerClusterToCluster(r)
	registerC2CMixedVersions(r)
	registerClusterReplicationResilience(r)
	registerClusterReplicationDisconnect(r)
	registerConnectionLatencyTest(r)
	registerCopy(r)
	registerCopyFrom(r)
	registerCostFuzz(r)
	registerDecommission(r)
	registerDecommissionBench(r)
	registerDisaggRebalance(r)
	registerDiskFull(r)
	registerDiskStalledDetection(r)
	registerDiskStalledWALFailover(r)
	registerDjango(r)
	registerDrain(r)
	registerDrop(r)
	registerEncryption(r)
	registerFailover(r)
	registerFixtures(r)
	registerFollowerReads(r)
	registerGORM(r)
	registerGopg(r)
	registerGossip(r)
	registerHibernate(r, hibernateOpts)
	registerHibernate(r, hibernateSpatialOpts)
	registerHotSpotSplits(r)
	registerHTTPRestart(r)
	registerFISmokeTest(r)
	registerImportCancellation(r)
	registerImportDecommissioned(r)
	registerImportMixedVersions(r)
	registerImportNodeShutdown(r)
	registerImportTPCC(r)
	registerImportTPCH(r)
	registerInconsistency(r)
	registerIndexes(r)
	registerJasyncSQL(r)
	registerJepsen(r)
	registerJobs(r)
	registerJobsMixedVersions(r)
	registerKerberosConnectionStressTest(r)
	registerKV(r)
	registerKVBench(r)
	registerKVContention(r)
	registerKVGracefulDraining(r)
	registerKVQuiescenceDead(r)
	registerKVRangeLookups(r)
	registerKVScalability(r)
	registerKVSplits(r)
	registerKVRestartImpact(r)
	registerKVStopAndCopy(r)
	registerKnex(r)
	registerLOQRecovery(r)
	registerLargeRange(r)
	registerLDAPConnectionLatencyTest(r)
	registerLDAPConnectionScaleTest(r)
	registerLeasePreferences(r)
	registerLedger(r)
	registerLibPQ(r)
	registerLiquibase(r)
	registerLoadSplits(r)
	registerLogicalDataReplicationTests(r)
	registerLDRMixedVersions(r)
	registerMVCCGC(r)
	registerMultiStoreRemove(r)
	registerMultiTenantDistSQL(r)
	registerMultiTenantMultiregion(r)
	registerMultiTenantTPCH(r)
	registerMultiTenantUpgrade(r)
	registerMultiTenantSharedProcess(r)
	registerNetwork(r)
	registerBufferedLogging(r)
	registerNodeJSPostgres(r)
	registerNpgsql(r)
	registerPebbleWriteThroughput(r)
	registerPebbleYCSB(r)
	registerPgjdbc(r)
	registerPGRegress(r)
	registerPgx(r)
	registerPointTombstone(r)
	registerPop(r)
	registerProcessLock(r)
	registerPsycopg(r)
	registerPruneDanglingSnapshotsAndDisks(r)
	registerPTP(r)
	registerQueue(r)
	registerQuitTransfersLeases(r)
	registerRebalanceLoad(r)
	registerReplicaGC(r)
	registerRestart(r)
	registerRestore(r)
	registerRestoreNodeShutdown(r)
	registerOnlineRestorePerf(r)
	registerOnlineRestoreCorrectness(r)
	registerRoachmart(r)
	registerRoachtest(r)
	registerRubyPG(r)
	registerRustPostgres(r)
	registerSQLAlchemy(r)
	registerSQLSmith(r)
	registerSchemaChangeBulkIngest(r)
	registerSchemaChangeDuringKV(r)
	registerSchemaChangeDuringTPCC800(r)
	registerSchemaChangeIndexTPCC100(r)
	registerSchemaChangeIndexTPCC800(r)
	registerSchemaChangeInvertedIndex(r)
	registerSchemaChangeMixedVersions(r)
	registerDeclSchemaChangeCompatMixedVersions(r)
	registerSchemaChangeRandomLoad(r)
	registerLargeSchemaBackupRestores(r)
	registerLargeSchemaBenchmarks(r)
	registerScrubAllChecksTPCC(r)
	registerScrubIndexOnlyTPCC(r)
	registerSecondaryIndexesMultiVersionCluster(r)
	registerSequelize(r)
	registerSlowDrain(r)
	registerSysbench(r)
	registerTLP(r)
	registerTPCC(r)
	registerTPCDSVec(r)
	registerTPCE(r)
	registerTPCHBench(r)
	registerTPCHConcurrency(r)
	registerTPCHVec(r)
	registerTypeORM(r)
	registerUnoptimizedQueryOracle(r)
	registerValidateSystemSchemaAfterVersionUpgradeSeparateProcess(r)
	registerYCSB(r)
	registerDeclarativeSchemaChangerJobCompatibilityInMixedVersion(r)
	registerMultiRegionMixedVersion(r)
	registerMultiRegionSystemDatabase(r)
	registerSqlStatsMixedVersion(r)
	registerDbConsole(r)
	perturbation.RegisterTests(r)
}
