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

// Provision cloudant_activity_tracker_events resource instance
resource "ibm_cloudant_activity_tracker_events" "cloudant_activity_tracker_events_instance" {
  types = var.cloudant_activity_tracker_events_types
}

// Provision cloudant_capacity_throughput resource instance
resource "ibm_cloudant_capacity_throughput" "cloudant_capacity_throughput_instance" {
  blocks = var.cloudant_capacity_throughput_blocks
}

// Provision cloudant_cors resource instance
resource "ibm_cloudant_cors" "cloudant_cors_instance" {
  origins = var.cloudant_cors_config_origins
  allow_credentials = var.cloudant_cors_config_allow_credentials
  enable_cors = var.cloudant_enable_cors
}

// Create cloudant data source
data "ibm_cloudant" "cloudant" {
  name     = ibm_cloudant.cloudant.name
}

// Create cloudant_server_information data source
data "ibm_cloudant_server_information" "cloudant_server_information_instance" {
}

// Create cloudant_activity_tracker_events data source
data "ibm_cloudant_activity_tracker_events" "cloudant_activity_tracker_events_instance" {
}

// Create cloudant_capacity_throughput data source
data "ibm_cloudant_capacity_throughput" "cloudant_capacity_throughput_instance" {
}

// Create cloudant_cors data source
data "ibm_cloudant_cors" "cloudant_cors_instance" {
}
