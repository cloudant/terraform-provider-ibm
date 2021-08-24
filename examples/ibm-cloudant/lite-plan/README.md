# IBM Cloudant example for Lite plan

This example shows how to create a Lite plan Cloudant instance.

To run, configure your IBM Cloudant provider

Running the example

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

cloudant resource:

```hcl
resource "ibm_cloudant" "cloudant" {
  name     = "test_lite_plan_cloudant"
  location = var.service_region
  plan     = "lite"
}
```

## Data sources

cloudant data source:

```hcl
data "ibm_cloudant" "cloudant" {
  name     = ibm_cloudant.cloudant.name
}
```
