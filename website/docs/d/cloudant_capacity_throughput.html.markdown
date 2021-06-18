---
layout: "ibm"
page_title: "IBM : cloudant_capacity_throughput"
description: |-
  Get information about cloudant_capacity_throughput
subcategory: "Cloudant"
---

# ibm\_cloudant_capacity_throughput

Provides a read-only data source for cloudant_capacity_throughput. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "cloudant_capacity_throughput" "cloudant_capacity_throughput" {
}
```

## Argument Reference

The following arguments are supported:


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the cloudant_capacity_throughput.
* `current` - Detailed information about provisioned throughput capacity. Nested `current` blocks have the following structure:
	* `throughput` - Schema for detailed information about throughput capacity with breakdown by specific throughput requests classes. Nested `throughput` blocks have the following structure:
		* `blocks` - A number of blocks of throughput units. A block consists of 100 reads/sec, 50 writes/sec, and 5 global queries/sec of provisioned throughput capacity.
		* `query` - Provisioned global queries capacity in operations per second.
		* `read` - Provisioned reads capacity in operations per second.
		* `write` - Provisioned writes capacity in operations per second.

