provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
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
  origins = var.cloudant_cors_origins
  allow_credentials = var.cloudant_cors_allow_credentials
  enable_cors = var.cloudant_cors_enable_cors
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
