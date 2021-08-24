provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region           = var.service_region
}

// Provision cloudant resource instance with Standard plan and dedicated hardware CRN
resource "ibm_cloudant" "cloudant" {
  // Required arguments:
  name     = "test_standard_plan_on_dedicated_hw_cloudant"
  location = var.service_region
  plan     = "standard"
  // Optional arguments:
  environment_crn = var.cloudant_cluster_information
}

// Create cloudant data source
data "ibm_cloudant" "cloudant" {
  name     = ibm_cloudant.cloudant.name
}
