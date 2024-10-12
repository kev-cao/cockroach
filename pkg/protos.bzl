# Code generated by generate-bazel-extra, DO NOT EDIT.

# This lists all the direct and indirect proto_library dependencies of
# //pkg/server/serverpb:serverpb_proto.
SERVER_PROTOS = [
    "//pkg/build:build_proto",
    "//pkg/clusterversion:clusterversion_proto",
    "//pkg/config/zonepb:zonepb_proto",
    "//pkg/geo/geopb:geopb_proto",
    "//pkg/gossip:gossip_proto",
    "//pkg/jobs/jobspb:jobspb_proto",
    "//pkg/kv/kvpb:kvpb_proto",
    "//pkg/kv/kvserver/concurrency/isolation:isolation_proto",
    "//pkg/kv/kvserver/concurrency/lock:lock_proto",
    "//pkg/kv/kvserver/kvflowcontrol/kvflowcontrolpb:kvflowcontrolpb_proto",
    "//pkg/kv/kvserver/kvserverpb:kvserverpb_proto",
    "//pkg/kv/kvserver/liveness/livenesspb:livenesspb_proto",
    "//pkg/kv/kvserver/loqrecovery/loqrecoverypb:loqrecoverypb_proto",
    "//pkg/kv/kvserver/readsummary/rspb:rspb_proto",
    "//pkg/multitenant/mtinfopb:mtinfopb_proto",
    "//pkg/multitenant/tenantcapabilities/tenantcapabilitiespb:tenantcapabilitiespb_proto",
    "//pkg/raft/raftpb:raftpb_proto",
    "//pkg/roachpb:roachpb_proto",
    "//pkg/rpc/rpcpb:rpcpb_proto",
    "//pkg/server/diagnostics/diagnosticspb:diagnosticspb_proto",
    "//pkg/server/serverpb:serverpb_proto",
    "//pkg/server/status/statuspb:statuspb_proto",
    "//pkg/settings:settings_proto",
    "//pkg/sql/appstatspb:appstatspb_proto",
    "//pkg/sql/catalog/catenumpb:catenumpb_proto",
    "//pkg/sql/catalog/catpb:catpb_proto",
    "//pkg/sql/catalog/descpb:descpb_proto",
    "//pkg/sql/catalog/fetchpb:fetchpb_proto",
    "//pkg/sql/contentionpb:contentionpb_proto",
    "//pkg/sql/lex:lex_proto",
    "//pkg/sql/schemachanger/scpb:scpb_proto",
    "//pkg/sql/sem/semenumpb:semenumpb_proto",
    "//pkg/sql/sessiondatapb:sessiondatapb_proto",
    "//pkg/sql/sqlstats/insights:insights_proto",
    "//pkg/sql/types:types_proto",
    "//pkg/storage/enginepb:enginepb_proto",
    "//pkg/ts/catalog:catalog_proto",
    "//pkg/ts/tspb:tspb_proto",
    "//pkg/util/admission/admissionpb:admissionpb_proto",
    "//pkg/util/duration:duration_proto",
    "//pkg/util/hlc:hlc_proto",
    "//pkg/util/log/logpb:logpb_proto",
    "//pkg/util/metric:metric_proto",
    "//pkg/util/timeutil/pgdate:pgdate_proto",
    "//pkg/util/tracing/tracingpb:tracingpb_proto",
    "//pkg/util:util_proto",
    "@com_github_cockroachdb_errors//errorspb:errorspb_proto",
    "@com_github_gogo_protobuf//gogoproto:gogo_proto",
    "@com_github_prometheus_client_model//io/prometheus/client:io_prometheus_client_proto",
    "@com_google_protobuf//:any_proto",
    "@com_google_protobuf//:descriptor_proto",
    "@com_google_protobuf//:duration_proto",
    "@com_google_protobuf//:timestamp_proto",
    "@go_googleapis//google/api:annotations_proto",
    "@go_googleapis//google/api:http_proto",
]

# This lists all the in-tree .proto files required to build serverpb_proto.
PROTO_FILES = [
    "//pkg/build:info.proto",
    "//pkg/clusterversion:cluster_version.proto",
    "//pkg/config/zonepb:zone.proto",
    "//pkg/geo/geopb:config.proto",
    "//pkg/geo/geopb:geopb.proto",
    "//pkg/gossip:gossip.proto",
    "//pkg/jobs/jobspb:jobs.proto",
    "//pkg/jobs/jobspb:schedule.proto",
    "//pkg/kv/kvpb:api.proto",
    "//pkg/kv/kvpb:errors.proto",
    "//pkg/kv/kvserver/concurrency/isolation:levels.proto",
    "//pkg/kv/kvserver/concurrency/lock:lock_waiter.proto",
    "//pkg/kv/kvserver/concurrency/lock:locking.proto",
    "//pkg/kv/kvserver/kvflowcontrol/kvflowcontrolpb:kvflowcontrol.proto",
    "//pkg/kv/kvserver/kvserverpb:internal_raft.proto",
    "//pkg/kv/kvserver/kvserverpb:lease_status.proto",
    "//pkg/kv/kvserver/kvserverpb:proposer_kv.proto",
    "//pkg/kv/kvserver/kvserverpb:raft.proto",
    "//pkg/kv/kvserver/kvserverpb:range_log.proto",
    "//pkg/kv/kvserver/kvserverpb:state.proto",
    "//pkg/kv/kvserver/liveness/livenesspb:liveness.proto",
    "//pkg/kv/kvserver/loqrecovery/loqrecoverypb:recovery.proto",
    "//pkg/kv/kvserver/readsummary/rspb:summary.proto",
    "//pkg/multitenant/mtinfopb:info.proto",
    "//pkg/multitenant/tenantcapabilities/tenantcapabilitiespb:capabilities.proto",
    "//pkg/raft/raftpb:raft.proto",
    "//pkg/roachpb:data.proto",
    "//pkg/roachpb:index_usage_stats.proto",
    "//pkg/roachpb:internal.proto",
    "//pkg/roachpb:io-formats.proto",
    "//pkg/roachpb:metadata.proto",
    "//pkg/roachpb:span_config.proto",
    "//pkg/roachpb:span_stats.proto",
    "//pkg/rpc/rpcpb:rpc.proto",
    "//pkg/server/diagnostics/diagnosticspb:diagnostics.proto",
    "//pkg/server/serverpb:admin.proto",
    "//pkg/server/serverpb:authentication.proto",
    "//pkg/server/serverpb:index_recommendations.proto",
    "//pkg/server/serverpb:init.proto",
    "//pkg/server/serverpb:migration.proto",
    "//pkg/server/serverpb:status.proto",
    "//pkg/server/status/statuspb:status.proto",
    "//pkg/settings:encoding.proto",
    "//pkg/sql/appstatspb:app_stats.proto",
    "//pkg/sql/catalog/catenumpb:encoded_datum.proto",
    "//pkg/sql/catalog/catenumpb:index.proto",
    "//pkg/sql/catalog/catpb:catalog.proto",
    "//pkg/sql/catalog/catpb:enum.proto",
    "//pkg/sql/catalog/catpb:function.proto",
    "//pkg/sql/catalog/catpb:privilege.proto",
    "//pkg/sql/catalog/descpb:join_type.proto",
    "//pkg/sql/catalog/descpb:lease.proto",
    "//pkg/sql/catalog/descpb:locking.proto",
    "//pkg/sql/catalog/descpb:structured.proto",
    "//pkg/sql/catalog/fetchpb:index_fetch.proto",
    "//pkg/sql/contentionpb:contention.proto",
    "//pkg/sql/lex:encode.proto",
    "//pkg/sql/schemachanger/scpb:elements.proto",
    "//pkg/sql/schemachanger/scpb:scpb.proto",
    "//pkg/sql/sem/semenumpb:constraint.proto",
    "//pkg/sql/sem/semenumpb:trigger.proto",
    "//pkg/sql/sessiondatapb:local_only_session_data.proto",
    "//pkg/sql/sessiondatapb:session_data.proto",
    "//pkg/sql/sessiondatapb:session_migration.proto",
    "//pkg/sql/sessiondatapb:session_revival_token.proto",
    "//pkg/sql/sqlstats/insights:insights.proto",
    "//pkg/sql/types:types.proto",
    "//pkg/storage/enginepb:engine.proto",
    "//pkg/storage/enginepb:file_registry.proto",
    "//pkg/storage/enginepb:mvcc.proto",
    "//pkg/storage/enginepb:mvcc3.proto",
    "//pkg/storage/enginepb:rocksdb.proto",
    "//pkg/ts/catalog:chart_catalog.proto",
    "//pkg/ts/tspb:timeseries.proto",
    "//pkg/util/admission/admissionpb:admission_stats.proto",
    "//pkg/util/admission/admissionpb:io_threshold.proto",
    "//pkg/util/duration:duration.proto",
    "//pkg/util/hlc:legacy_timestamp.proto",
    "//pkg/util/hlc:timestamp.proto",
    "//pkg/util/log/logpb:event.proto",
    "//pkg/util/log/logpb:log.proto",
    "//pkg/util/metric:metric.proto",
    "//pkg/util/timeutil/pgdate:pgdate.proto",
    "//pkg/util/tracing/tracingpb:recorded_span.proto",
    "//pkg/util/tracing/tracingpb:tracing.proto",
    "//pkg/util:unresolved_addr.proto",
    "@com_github_cockroachdb_errors//errorspb:errors.proto",
    "@com_github_cockroachdb_errors//errorspb:hintdetail.proto",
    "@com_github_cockroachdb_errors//errorspb:markers.proto",
    "@com_github_cockroachdb_errors//errorspb:tags.proto",
    "@com_github_cockroachdb_errors//errorspb:testing.proto",
    "@com_github_prometheus_client_model//io/prometheus/client:metrics.proto",
]
