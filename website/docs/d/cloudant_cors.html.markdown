---
layout: "ibm"
page_title: "IBM : cloudant_cors"
description: |-
  Get information about cloudant_cors
subcategory: "Cloudant"
---

# ibm\_cloudant_cors

Provides a read-only data source for cloudant_cors. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "cloudant_cors" "cloudant_cors" {
}
```

## Argument Reference

The following arguments are supported:


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the cloudant_cors.
* `allow_credentials` - Boolean value to allow authentication credentials. If set to true, browser requests must be done by using withCredentials = true.

* `enable_cors` - Boolean value to turn CORS on and off.

* `origins` - An array of strings that contain allowed origin domains. You have to specify the full URL including the protocol. It is recommended that only the HTTPS protocol is used. Subdomains count as separate domains, so you have to specify all subdomains used.

