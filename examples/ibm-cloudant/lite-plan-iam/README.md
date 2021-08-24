# IBM Cloudant example for Lite plan with IAM credentials

This example shows how to create a Lite plan Cloudant instance with Viewer IAM credential.

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
  name     = "test_lite_plan_iam_cloudant"
  location = var.service_region
  plan     = "lite"
}
```

resourceKey resource to create a new IAM credentials with Viewer role:

```hcl
resource "ibm_resource_key" "resourceKey" {
  name                 = "myCredentials"
  role                 = "Viewer"
  resource_instance_id = data.ibm_cloudant.cloudant.id
}
```

## Data sources

cloudant data source:

```hcl
data "ibm_cloudant" "cloudant" {
  name     = ibm_cloudant.cloudant.name
}
```
