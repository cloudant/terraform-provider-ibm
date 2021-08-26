# Examples for `ibm_cloudant`

Examples in the subfolders illustrate how to use the `ibm_cloudant`.

These types of resources are supported:

* ibm_cloudant

Each example creates an IBM Cloudant resource instance called `cloudant`.

## Examples

Examples can be found in the subfolders along with the instuctions how to run them.

- [Lite plan](lite-plan)
- [Lite plan with legacy credentials](lite-plan-legacy)
- [Lite plan with IAM credentials](lite-plan-iam)
- [Standard plan with custom capacity](standard-plan)
- [Standard plan with data event tracking](standard-plan-with-data-events)
- [Standard plan on dedicated hardware](standard-plan-on-dedicated-hw)

## Assumptions

## Notes

1. With `Lite` plan `capacity` can be set no more than 1 throughput blocks.
1. `parameters` can overwrite the previously set arguments named the same way.
1. With [`Standard` plan on dedicated hardware](standard-plan-on-dedicated-hw) the hardware must be ordered separately and provisioning should be completed before using Terraform on it

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud_api_key | IBM Cloud API key. | `string` | true |
| name | Service instance name. | `string` | true |
| location | Region in which service has to be provisioned. | `string` | true |
| plan | The plan type of the service. | string | true |
| id | The ID of the service. | string | false |
| types | An array of event types that are being sent to IBM Cloud Activity Tracker for the IBM Cloudant instance. Allowable values: management, data. "management" is a required element of this array. | `list(string)` | false |
| blocks | A number of blocks of throughput units. A block consists of 100 reads/sec, 50 writes/sec, and 5 global queries/sec of provisioned throughput capacity. | `number` | false |
| cluster_information | The actual physical location of the Dedicated Hardware plan instance. | `string` | false
| cors_config.origins | An array of strings that contain allowed origin domains. You have to specify the full URL including the protocol. It is recommended that only the HTTPS protocol is used. Subdomains count as separate domains, so you have to specify all subdomains used. | `list(string)` | false |
| cors_config.allow_credentials | Boolean value to allow authentication credentials. If set to true, browser requests must be done by using withCredentials = true. | `bool` | false |
| enable_cors | Boolean value to turn CORS on and off. | `bool` | false |
| environment_crn | CRN of the IBM Cloudant Dedicated Hardware plan instance. | `string` | false |
| legacy_credentials | Use both legacy credentials and IAM for authentication. | `bool` | false |
| include_data_events | Include data event types in events sent to IBM Cloud Activity Tracker with LogDNA for the IBM Cloudant instance. By default only emitted events are of \"management\" type. | `bool` | false
| instance_crn | CRN of the Key Protect instance housing the encryption key for BYOK. | `string` | false |
| parameters | Arbitrary parameters to pass. Must be a JSON object. | `map(string)` | false |
| resource_group_id | The resource group id. | `string` | false |
| service_endpoints | Types of the service endpoints. Possible values are 'public', 'private', 'public-and-private'. | `string` | false |
| tags | | `set(string)` | false |
| timeouts.create<br>timeouts.update<br>timeouts.delete | The operation of the IBM Cloudant instance is considered failed if no response received for the given timeout. | `string` | false |

## Outputs

| Name | Description |
|------|-------------|
| ibm_cloudant | `ibm_cloudant` terraform resource instance |
