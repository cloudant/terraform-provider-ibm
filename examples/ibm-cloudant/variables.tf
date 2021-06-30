// Resource arguments for ibmcloud_api_key
variable "ibmcloud_api_key" {
  description = "IBM Cloud API key."
  type        = string
}

// Resource arguments for service_name
variable "service_name" {
  description = "Service instance name."
  type        = string
}

// Resource arguments for service_region
variable "service_location" {
  description = "Region in which service has to be provisioned."
  type        = string
  default     = "us-south"
}

// Resource arguments for service_plan
variable "service_plan" {
  description = "The plan type of the service."
  type        = string
  default     = "lite"
}

// Resource arguments for service_id
variable "service_id" {
  description = "The ID of the service."
  type        = string
}

// Resource arguments for cloudant_activity_tracker_events
variable "cloudant_activity_tracker_events_types" {
  description = "An array of event types that are being sent to IBM Cloud Activity Tracker for the IBM Cloudant instance. Allowable values: management, data. \"management\" is a required element of this array."
  type        = list(string)
  default     = [ "management" ]
}

// Resource arguments for cloudant_capacity_throughput
variable "cloudant_capacity_throughput_blocks" {
  description = "A number of blocks of throughput units. A block consists of 100 reads/sec, 50 writes/sec, and 5 global queries/sec of provisioned throughput capacity."
  type        = number
  default     = 1
}

// Resource arguments for cloudant_cluster_information
variable "cloudant_cluster_information" {
  description = "The actual physical location of the Dedicated Hardware plan instance."
  type = string
}

// Resource arguments for cloudant_cors
variable "cloudant_cors_config_origins" {
  description = "An array of strings that contain allowed origin domains. You have to specify the full URL including the protocol. It is recommended that only the HTTPS protocol is used. Subdomains count as separate domains, so you have to specify all subdomains used."
  type        = list(string)
  default     = [ "https://example.com" ]
}
variable "cloudant_cors_config_allow_credentials" {
  description = "Boolean value to allow authentication credentials. If set to true, browser requests must be done by using withCredentials = true."
  type        = bool
  default     = false
}
variable "cloudant_enable_cors" {
  description = "Boolean value to turn CORS on and off."
  type        = bool
  default     = false
}

// Resource arguments for cloudant_environment_crn
variable "cloudant_environment_crn" {
  description = "CRN of the IBM Cloudant Dedicated Hardware plan instance."
  type = string
}

// Resource arguments for cloudant_hipaa
variable "cloudant_hipaa" {
  description = "Instance is HIPAA ready in US locations."
  type = bool
}

// Resource arguments for cloudant_legacy_credentials
variable "cloudant_legacy_credentials" {
  description = "Use both legacy credentials and IAM for authentication."
  type = bool
  default = false
}

// Resource arguments for cloudant_include_data_events
variable "cloudant_include_data_events" {
  description = "Include data event types in events sent to IBM Cloud Activity Tracker with LogDNA for the IBM Cloudant instance. By default only emitted events are of \"management\" type."
  type = bool
  default = false
}

// Resource arguments for cloudant_instance_crn
variable "cloudant_instance_crn" {
  description = "CRN of the Key Protect instance housing the encryption key for BYOK."
  type = string
}

// Resource arguments for cloudant_key_crn
variable "cloudant_key_crn" {
  description = "CRN of the encryption key that is stored in the Key Protect instance."
  type = string
}

// Resource arguments for cloudant_parameters
variable "cloudant_parameters" {
  description = "Arbitrary parameters to pass. Must be a JSON object."
  type = map(string)
}
// Resource arguments for cloudant_resource_group_id
variable "cloudant_resource_group_id" {
  description = "The resource group id."
  type = string
}

// Resource arguments for cloudant_service_endpoints
variable "cloudant_service_endpoints" {
  description = "Types of the service endpoints. Possible values are 'public', 'private', 'public-and-private'."
  type = string
  default = "public"
}

// Resource arguments for cloudant_tags
variable "cloudant_tags" {
  type = set(string)
}

// Resource arguments for timeout
variable "cloudant_timeout" {
  description = "The operation of the IBM Cloudant instance is considered failed if no response received for the given timeout."
  type = string
  default = "15m"
}

// Data source arguments for cloudant_server_information

// Data source arguments for cloudant_activity_tracker_events

// Data source arguments for cloudant_capacity_throughput

// Data source arguments for cloudant_cors
