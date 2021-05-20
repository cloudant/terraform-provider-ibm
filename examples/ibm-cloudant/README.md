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
resource "cloudant_activity_tracker_events" "cloudant_activity_tracker_events_instance" {
  types = var.cloudant_activity_tracker_events_types
}
```
cloudant_capacity_throughput resource:

```hcl
resource "cloudant_capacity_throughput" "cloudant_capacity_throughput_instance" {
  blocks = var.cloudant_capacity_throughput_blocks
}
```
cloudant_cors resource:

```hcl
resource "cloudant_cors" "cloudant_cors_instance" {
  origins = var.cloudant_cors_origins
  allow_credentials = var.cloudant_cors_allow_credentials
  enable_cors = var.cloudant_cors_enable_cors
}
```

## CloudantV1 Data sources

cloudant_server_information data source:

```hcl
data "cloudant_server_information" "cloudant_server_information_instance" {
}
```
cloudant_activity_tracker_events data source:

```hcl
data "cloudant_activity_tracker_events" "cloudant_activity_tracker_events_instance" {
}
```
cloudant_capacity_throughput data source:

```hcl
data "cloudant_capacity_throughput" "cloudant_capacity_throughput_instance" {
}
```
cloudant_cors data source:

```hcl
data "cloudant_cors" "cloudant_cors_instance" {
}
```

## Assumptions

1. TODO

## Notes

1. TODO

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
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| types | An array of event types that are being sent to IBM Cloud Activity Tracker for the IBM Cloudant instance. Allowable values: management, data. "management" is a required element of this array. | `list(string)` | true |
| blocks | A number of blocks of throughput units. A block consists of 100 reads/sec, 50 writes/sec, and 5 global queries/sec of provisioned throughput capacity. | `number` | true |
| origins | An array of strings that contain allowed origin domains. You have to specify the full URL including the protocol. It is recommended that only the HTTPS protocol is used. Subdomains count as separate domains, so you have to specify all subdomains used. | `list(string)` | true |
| allow_credentials | Boolean value to allow authentication credentials. If set to true, browser requests must be done by using withCredentials = true. | `bool` | false |
| enable_cors | Boolean value to turn CORS on and off. | `bool` | false |

## Outputs

| Name | Description |
|------|-------------|
| cloudant_activity_tracker_events | cloudant_activity_tracker_events object |
| cloudant_capacity_throughput | cloudant_capacity_throughput object |
| cloudant_cors | cloudant_cors object |
| cloudant_server_information | cloudant_server_information object |
| cloudant_activity_tracker_events | cloudant_activity_tracker_events object |
| cloudant_capacity_throughput | cloudant_capacity_throughput object |
| cloudant_cors | cloudant_cors object |
