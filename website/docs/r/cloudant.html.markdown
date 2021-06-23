---
layout: "ibm"
page_title: "IBM : ibm_cloudant"
description: |-
  Manages Cloudant instance.
subcategory: "Cloudant"
---

# ibm_cloudant

Provides a resource for IBM Cloudant. This allows an IBM Cloudant service instance to be created, updated, or deleted.
For more information, about how to use this service, see the official [Cloudant documentation](https://cloud.ibm.com/docs/Cloudant?topic=Cloudant-getting-started-with-cloudant).

## Example Usage

```hcl
resource "ibm_cloudant" "cloudant" {
  name     = "cloudant-service-name"
  location = "us-south"
  plan     = "standard"

  legacy_credentials  = true
  include_data_events = false
  capacity            = 1
  enable_cors         = true

  cors_config {
    allow_credentials = false
    origins           = ["https://example.com"]
  }

  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
}
```

## Timeouts

ibm_cloudant provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html#operation-timeouts) 
configuration options:

* `create` - (Default 10 minutes) The creation of the IBM Cloudant instance is considered failed if no response received.
* `delete` - (Default 10 minutes) The update of the IBM Cloudant instance is considered failed if no response received.
* `update` - (Default 10 minutes) The deletion of the IBM Cloudant instance is considered failed if no response received.

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) A name for the resource instance.
* `location` - (Required, Forces new resource, String) Target location or environment to create the resource instance.
* `plan` - (Required, String) The plan type of the service.
* `capacity` - (Optional, Number) A number of blocks of throughput units. A block consists of 100 reads/sec, 50 writes/sec, and 5 global queries/sec of provisioned throughput capacity.
    * Constraints: 
      * The default value is `1`.
      * Setting capacity is not supported for lite plan.
* `cluster_location` - (Optional, Forces new resource, String) The actual physical location of the Dedicated Hardware plan instance.
* `cors_config` - (Optional, Block List) Configuration for CORS. (see [below for nested arguments](#nestedblock--cors_config))
  * Constraints: 
    * The minimum length is `1` item.
    * Can conflict with `enable_cors` if it is set to `false`. In this case the CORS customization is not allowed.
* `enable_cors` - (Optional, Boolean) Boolean value to turn CORS on and off.
  * Constraints: 
    * The default value is `true`.
    * If it is set to `false`, then customizing `cors_config` is not allowed.
* `environment_crn` - (Optional, Forces new resource, String) CRN of the IBM Cloudant Dedicated Hardware plan instance.
* `hipaa` - (Optional, Forces new resource, Boolean) Instance is HIPAA ready in US locations.
  * Constraints: The default value is `false`.
* `include_data_events` - (Optional, Boolean) Include data event types in events sent to IBM Cloud Activity Tracker with LogDNA for the IBM Cloudant instance. By default only emitted events are of "management" type.
  * Constraints: The default value is `false`.
* `kms_instance_crn` - (Optional, Forces new resource, String) CRN of the Key Protect instance housing the encryption key for BYOK.
* `kms_key_crn` - (Optional, Forces new resource, String) CRN of the encryption key that is stored in the Key Protect instance.
* `legacy_credentials` - (Optional, Forces new resource, Boolean) Use both legacy credentials and IAM for authentication.
  * Constraints: The default value is `false`.
* `parameters` - (Optional, Forces new resource, Map) Arbitrary parameters to pass. Must be a JSON object.
* `resource_group_id` - (Optional, Forces new resource, String) The resource group id.
* `service_endpoints` - (Optional, String) Types of the service endpoints. Possible values are 'public', 'private', 'public-and-private'.
* `tags` - (Optional, Set of String) Tags associated with the instance.

<a id="nestedblock--cors_config"></a>
### Nested Arguments for `cors_config`

* `origins` - (Required, List of String) An array of strings that contain allowed origin domains. You have to specify the full URL including the protocol. It is recommended that only the HTTPS protocol is used. Subdomains count as separate domains, so you have to specify all subdomains used.
* `allow_credentials` - (Optional, Boolean) Boolean value to allow authentication credentials. If set to true, browser requests must be done by using withCredentials = true.
  * Constraints: The default value is `true`.

## Attribute Reference

In addition to all arguments above, you can access the following attribute references after your resource is created.:

* `account_id` - (String) An alpha-numeric value identifying the account ID.
* `allow_cleanup` - (Boolean) A boolean that dictates if the resource instance should be deleted (cleaned up) during the processing of a region instance delete call.
* `created_at` - (String) The date when the instance was created.
* `created_by` - (String) The subject who created the instance.
* `crn` - (String) CRN of resource instance.
* `dashboard_url` - (String) Dashboard URL to access resource.
* `deleted_at` - (String) The date when the instance was deleted.
* `deleted_by` - (String) The subject who deleted the instance.
* `extensions` - (Map) The extended metadata as a map associated with the resource instance.
* `guid` - (String) Guid of resource instance.
* `id` - (Optional, String) The unique identifier of the new Cloudant resource.
* `last_operation` - (Map) The status of the last operation requested on the instance.
* `locked` - (Boolean) A boolean that dictates if the resource instance should be deleted (cleaned up) during the processing of a region instance delete call.
* `plan_history` - (List of Object) The plan history of the instance.
* `resource_aliases_url` - (String) The relative path to the resource aliases for the instance.
* `resource_bindings_url` - (String) The relative path to the resource bindings for the instance.
* `resource_controller_url` - (String) The URL of the IBM Cloud dashboard that can be used to explore and view details about the resource.
* `resource_crn` - (String) The crn of the resource.
* `resource_group_crn` - (String) The long ID (full CRN) of the resource group.
* `resource_group_name` - (String) The resource group name in which resource is provisioned.
* `resource_id` - (String) The unique ID of the offering.
* `resource_keys_url` - (String) The relative path to the resource keys for the instance.
* `resource_name` - (String) The name of the resource.
* `resource_plan_id` - (String) The unique ID of the plan associated with the offering.
* `resource_status` - (String) The status of the resource.
* `restored_at` - (String) The date when the instance under reclamation was restored.
* `restored_by` - (String) The subject who restored the instance back from reclamation.
* `scheduled_reclaim_at` - (String) The date when the instance was scheduled for reclamation.
* `scheduled_reclaim_by` - (String) The subject who initiated the instance reclamation.
* `service` - (String) The service type of the instance.
* `state` - (String) The current state of the instance.
* `status` - (String) Status of resource instance.
* `sub_type` - (String) The sub-type of instance, e.g. cfaas .
* `target_crn` - (String) The full deployment CRN as defined in the global catalog.
* `throughput` - (Map of Number) Schema for detailed information about throughput capacity with breakdown by specific throughput requests classes.
* `type` - (String) The type of the instance, e.g. service_instance.
* `update_at` - (String) The date when the instance was last updated.
* `update_by` - (String) The subject who updated the instance.

## Import

You can import the `ibm_cloudant` resource by using `crn`. An array of event types that are being sent to IBM Cloud Activity Tracker for the IBM Cloudant instance. Allowable values: management, data. "management" is a required element of this array.

### Syntax

```hcl-terraform
$ terraform import ibm_cloudant.mycloudant <crn>
```

### Example
```hcl
$ terraform import ibm_cloudant.mycloudant "crn:v1:bluemix:public:cloudantnosqldb:us-south:a/4ea2849a2d3401ed1e457268541966ea:79262bd4-4067-4783-b5ce-b1dba84ff8c4::"
```
