# Table: aws_route53_delegation_sets

This table shows data for Route53 Delegation Sets.

https://docs.aws.amazon.com/Route53/latest/APIReference/API_DelegationSet.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn (PK)|`utf8`|
|name_servers|`list<item: utf8, nullable>`|
|caller_reference|`utf8`|
|id|`utf8`|