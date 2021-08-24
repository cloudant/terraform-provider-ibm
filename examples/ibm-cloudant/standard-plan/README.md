# IBM Cloudant example for Standard plan with custom capacity

This example shows how to create a Standard plan Cloudant instance with 2 capacity throughput blocks.

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

cloudant resource with 2 capacity throughput blocks:

```hcl
resource "ibm_cloudant" "cloudant" {
  name     = "test_standard_plan_cloudant"
  location = var.service_region
  plan     = "standard"
  capacity = 2
}
```

## Data sources

cloudant data source:

```hcl
data "ibm_cloudant" "cloudant" {
  name     = ibm_cloudant.cloudant.name
}
```
