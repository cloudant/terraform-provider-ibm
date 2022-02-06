---
layout: "ibm"
page_title: "IBM : cloudant_database"
description: |-
  Manages cloudant_database.
subcategory: "Cloudant"
---

# ibm\_cloudant_database

Provides a resource for cloudant_database. This allows cloudant_database to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cloudant_database" "cloudant_database" {
  instance_crn  = var.instance_crn
  db            = var.db_name
}  
```

## Argument Reference

The following arguments are supported:

* `instance_crn` - (Required, string) Path parameter to specify the cloudant instance CRN.
* `db` - (Required, Forces new resource, string) Path parameter to specify the database name.
* `partitioned` - (Optional, Forces new resource, bool) Query parameter to specify whether to enable database partitions when creating a database.
  * Constraints: The default value is `false`.
* `q` - (Optional, Forces new resource, int) The number of shards in the database. Each shard is a partition of the hash value range. Default is 8, unless overridden in the `cluster config`.
  * Constraints: The minimum value is `1`.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the cloudant_database.

## Import

You can import the `cloudant_database` resource by using `ID`.
The `ID` property can be formed from `instance_crn`, and `db` in the following format:

```
<instance_crn>/<db>
```
* `instance_crn`: A string. Path parameter to specify the cloudant instance CRN.
* `db`: A string. Path parameter to specify the database name.

```
$ terraform import ibm_cloudant_database.cloudant_database <instance_crn>/<db>
```
