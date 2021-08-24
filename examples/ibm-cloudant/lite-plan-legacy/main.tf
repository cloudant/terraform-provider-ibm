provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region           = var.service_region
}

// Provision cloudant resource instance with Lite plan
resource "ibm_cloudant" "cloudant" {
  // Required arguments:
  name     = "test_lite_plan_legacy_cloudant"
  location = var.service_region
  plan     = "lite"
  // Optional arguments:
  legacy_credentials = true // Use both legacy credentials and IAM for authentication.
}

// Create cloudant data source
data "ibm_cloudant" "cloudant" {
  name     = ibm_cloudant.cloudant.name
}
