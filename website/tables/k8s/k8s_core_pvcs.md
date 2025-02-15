# Table: k8s_core_pvcs

This table shows data for Kubernetes (K8s) Core Persistent Volume Claims (PVCs).

The primary key for this table is **uid**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|context|`utf8`|
|kind|`utf8`|
|api_version|`utf8`|
|name|`utf8`|
|namespace|`utf8`|
|uid (PK)|`utf8`|
|resource_version|`utf8`|
|generation|`int64`|
|deletion_grace_period_seconds|`int64`|
|labels|`json`|
|annotations|`json`|
|owner_references|`json`|
|finalizers|`list<item: utf8, nullable>`|
|spec_access_modes|`list<item: utf8, nullable>`|
|spec_selector|`json`|
|spec_resources|`json`|
|spec_volume_name|`utf8`|
|spec_storage_class_name|`utf8`|
|spec_volume_mode|`utf8`|
|spec_data_source|`json`|
|spec_data_source_ref|`json`|
|status_phase|`utf8`|
|status_access_modes|`list<item: utf8, nullable>`|
|status_capacity|`json`|
|status_conditions|`json`|
|status_allocated_resources|`json`|
|status_resize_status|`utf8`|