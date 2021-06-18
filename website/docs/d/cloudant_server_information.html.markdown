---
layout: "ibm"
page_title: "IBM : cloudant_server_information"
description: |-
  Get information about cloudant_server_information
subcategory: "Cloudant"
---

# ibm\_cloudant_server_information

Provides a read-only data source for cloudant_server_information. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "cloudant_server_information" "cloudant_server_information" {
}
```

## Argument Reference

The following arguments are supported:


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the cloudant_server_information.
* `features` - List of enabled optional features.

* `vendor` - Schema for server vendor information. Nested `vendor` blocks have the following structure:
	* `version` - Vendor version.

* `features_flags` - List of feature flags.

