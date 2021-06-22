# Example for CloudantV1

This example illustrates how to use the CloudantV1

These types of resources are supported:

* cloudant_activity_tracker_events
* cloudant_capacity_throughput
* cloudant_cors

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## CloudantV1 resources

cloudant_activity_tracker_events resource:

```hcl
resource "ibm_cloudant_activity_tracker_events" "cloudant_activity_tracker_events_instance" {
  types = var.cloudant_activity_tracker_events_types
}
```
cloudant_capacity_throughput resource:

```hcl
resource "ibm_cloudant_capacity_throughput" "cloudant_capacity_throughput_instance" {
  blocks = var.cloudant_capacity_throughput_blocks
}
```
cloudant_cors resource:

```hcl
resource "ibm_cloudant_cors" "cloudant_cors_instance" {
  origins = var.cloudant_cors_config_origins
  allow_credentials = var.cloudant_cors_config_allow_credentials
  enable_cors = var.cloudant_enable_cors
}
```

## CloudantV1 Data sources

cloudant_server_information data source:

```hcl
data "ibm_cloudant_server_information" "cloudant_server_information_instance" {
}
```
cloudant_activity_tracker_events data source:

```hcl
data "ibm_cloudant_activity_tracker_events" "cloudant_activity_tracker_events_instance" {
}
```
cloudant_capacity_throughput data source:

```hcl
data "ibm_cloudant_capacity_throughput" "cloudant_capacity_throughput_instance" {
}
```
cloudant_cors data source:

```hcl
data "ibm_cloudant_cors" "cloudant_cors_instance" {
}
```

## Assumptions


## Notes

1. With `Lite` plan capacity_throughput_blocks can be set no more than 1.
1. `parameters` can overwrite the previously set arguments named the same way.

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
| hipaa | Instance is HIPAA ready in US locations. | `bool` | false |
| legacy_credentials | Use both legacy credentials and IAM for authentication. | `bool` | false |
| include_data_events | Include data event types in events sent to IBM Cloud Activity Tracker with LogDNA for the IBM Cloudant instance. By default only emitted events are of \"management\" type. | `bool` | false
| instance_crn | CRN of the Key Protect instance housing the encryption key for BYOK. | `string` | false |
| key_crn | CRN of the encryption key that is stored in the Key Protect instance. | `string` | false |
| parameters | Arbitrary parameters to pass. Must be a JSON object. | `map(string)` | false |
| resource_group_id | The resource group id. | `string` | false |
| service_endpoints | Types of the service endpoints. Possible values are 'public', 'private', 'public-and-private'. | `string` | false |
| tags | | `set(string)` | false |
| timeouts.create<br>timeouts.update<br>timeouts.delete | The operation of the IBM Cloudant instance is considered failed if no response received for the given timeout. | `string` | false |

## Outputs

| Name | Description |
|------|-------------|
| cloudant_activity_tracker_events | cloudant_activity_tracker_events object |
| cloudant_capacity_throughput | cloudant_capacity_throughput object |
| cloudant_cors | cloudant_cors object |
| cloudant_server_information | cloudant_server_information object |
