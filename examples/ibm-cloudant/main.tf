provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision cloudant resource instance
resource "ibm_cloudant" "cloudant" {
  // Required arguments:
  name     = var.service_name
  location = var.service_location
  plan     = var.service_plan

  // Optional arguments:
  capacity = var.cloudant_capacity_throughput_blocks

  enable_cors = var.cloudant_enable_cors
  cors_config {
    allow_credentials = var.cloudant_cors_config_allow_credentials
    origins           = var.cloudant_cors_config_origins
  }

  legacy_credentials = var.cloudant_legacy_credentials
  include_data_events = var.cloudant_include_data_events

  timeouts {
    create = var.cloudant_timeout
    update = var.cloudant_timeout
    delete = var.cloudant_timeout
  }
}

// Create cloudant data source
data "ibm_cloudant" "cloudant" {
  name     = ibm_cloudant.cloudant.name
}
