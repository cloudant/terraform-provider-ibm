provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region           = var.service_region
}

// Provision cloudant resource instance with Standard plan including data events
resource "ibm_cloudant" "cloudant" {
  // Required arguments:
  name     = "test_standard_plan_with_data_events_cloudant"
  location = var.service_region
  plan     = "standard"
  // Optional arguments:
  include_data_events = true
}

// Create cloudant data source
data "ibm_cloudant" "cloudant" {
  name     = ibm_cloudant.cloudant.name
}

// Provision activity tracker as a service in the same resource group as the cloudant resource instance
// See also https://registry.terraform.io/modules/terraform-ibm-modules/observability/ibm/latest/examples/activity-tracker-logdna-instance
module "activity_tracker_instance" {
  source = "terraform-ibm-modules/observability/ibm//modules/activity-tracker-logdna"

  service_name      = data.ibm_cloudant.cloudant.name
  plan              = "7-day"
  region            = var.service_region
  resource_group_id = data.ibm_cloudant.cloudant.id
}
