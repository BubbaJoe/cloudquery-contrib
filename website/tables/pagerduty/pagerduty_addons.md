# Table: pagerduty_addons

This table shows data for PagerDuty Addons.

https://developer.pagerduty.com/api-reference/e58b140202a57-list-installed-add-ons

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|type|`utf8`|
|summary|`utf8`|
|self|`utf8`|
|html_url|`utf8`|
|name|`utf8`|
|src|`utf8`|
|services|`json`|