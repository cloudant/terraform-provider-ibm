---
layout: "ibm"
page_title: "IBM : cloudant_capacity_throughput"
description: |-
  Manages cloudant_capacity_throughput.
subcategory: "Cloudant"
---

# ibm\_cloudant_capacity_throughput

Provides a resource for cloudant_capacity_throughput. This allows cloudant_capacity_throughput to be created, updated and deleted.

## Example Usage

```hcl
resource "cloudant_capacity_throughput" "cloudant_capacity_throughput" {
  blocks = 0
}
```

## Argument Reference

The following arguments are supported:

* `blocks` - (Required, Forces new resource, int) A number of blocks of throughput units. A block consists of 100 reads/sec, 50 writes/sec, and 5 global queries/sec of provisioned throughput capacity.
  * Constraints: The minimum value is `0`.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the cloudant_capacity_throughput.

## Import

You can import the `cloudant_capacity_throughput` resource by using `current`.
The `current` property can be formed from 

```

```

```
$ terraform import cloudant_capacity_throughput.cloudant_capacity_throughput 
```
