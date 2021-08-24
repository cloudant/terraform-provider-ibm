# IBM Cloudant example for Standard plan with data event tracking

This example shows how to create a Standard plan Cloudant instance with data event tracking.

To run, configure your IBM Cloudant provider

Running the example

For init the Activity Tracker module

```sh
terraform init
```

For planning phase

```sh
terraform plan
```

For apply phase

```sh
terraform apply
```

For destroy

```sh
terraform destroy
```

## Resources

cloudant resource with dedicated hardware CRN:

```hcl
resource "ibm_cloudant" "cloudant" {
  name     = "test_standard_plan_with_data_events_cloudant"
  location = var.service_region
  plan     = "standard"
}
```

## Data sources

cloudant data source:

```hcl
data "ibm_cloudant" "cloudant" {
  name     = ibm_cloudant.cloudant.name
}
```

## Modules

activity_tracker_instance module:

```hcl
module "activity_tracker_instance" {
  source = "terraform-ibm-modules/observability/ibm//modules/activity-tracker-logdna"

  service_name      = data.ibm_cloudant.cloudant.name
  plan              = "7-day"
  region            = var.service_region
  resource_group_id = data.ibm_cloudant.cloudant.id
}
```
