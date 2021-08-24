# IBM Cloudant example for Lite plan with legacy credentials

This example shows how to create a Lite plan Cloudant instance with legacy credentials.

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

cloudant resource with enabled legacy credentials:

```hcl
resource "ibm_cloudant" "cloudant" {
  name     = "test_lite_plan_legacy_cloudant"
  location = var.service_region
  plan     = "lite"
  legacy_credentials = true
}
```

## Data sources

cloudant data source:

```hcl
data "ibm_cloudant" "cloudant" {
  name     = ibm_cloudant.cloudant.name
}
```
