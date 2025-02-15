# Table: aws_ssm_inventory_schemas

This table shows data for AWS Systems Manager (SSM) Inventory Schemas.

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_InventoryItemSchema.html

The composite primary key for this table is (**account_id**, **region**, **type_name**, **version**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|type_name (PK)|`utf8`|
|version (PK)|`utf8`|
|attributes|`json`|
|display_name|`utf8`|