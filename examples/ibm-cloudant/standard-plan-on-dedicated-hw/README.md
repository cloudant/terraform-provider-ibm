# IBM Cloudant example for Standard plan on dedicated hardware

This example shows how to create a Standard plan Cloudant instance with dedicated hardware CRN.

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

cloudant resource with dedicated hardware CRN:

```hcl
resource "ibm_cloudant" "cloudant" {
  name     = "test_standard_plan_on_dedicated_hw_cloudant"
  location = var.service_region
  plan     = "standard"
  environment_crn = var.cloudant_cluster_information
}
```

## Data sources

cloudant data source:

```hcl
data "ibm_cloudant" "cloudant" {
  name     = ibm_cloudant.cloudant.name
}
```
