---
layout: "ibm"
page_title: "IBM : cloudant_cors"
description: |-
  Manages cloudant_cors.
subcategory: "Cloudant"
---

# ibm\_cloudant_cors

Provides a resource for cloudant_cors. This allows cloudant_cors to be created, updated and deleted.

## Example Usage

```hcl
resource "cloudant_cors" "cloudant_cors" {
}
```

## Argument Reference

The following arguments are supported:

* `origins` - (Required, Forces new resource, List) An array of strings that contain allowed origin domains. You have to specify the full URL including the protocol. It is recommended that only the HTTPS protocol is used. Subdomains count as separate domains, so you have to specify all subdomains used.
* `allow_credentials` - (Optional, Forces new resource, bool) Boolean value to allow authentication credentials. If set to true, browser requests must be done by using withCredentials = true.
  * Constraints: The default value is `true`.
* `enable_cors` - (Optional, Forces new resource, bool) Boolean value to turn CORS on and off.
  * Constraints: The default value is `true`.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the cloudant_cors.

## Import

You can import the `cloudant_cors` resource by using `enable_cors`. Boolean value to turn CORS on and off.

```
$ terraform import cloudant_cors.cloudant_cors <enable_cors>
```
