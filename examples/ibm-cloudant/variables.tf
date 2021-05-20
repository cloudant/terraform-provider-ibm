variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
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
  default     = 0
}

// Resource arguments for cloudant_cors
variable "cloudant_cors_origins" {
  description = "An array of strings that contain allowed origin domains. You have to specify the full URL including the protocol. It is recommended that only the HTTPS protocol is used. Subdomains count as separate domains, so you have to specify all subdomains used."
  type        = list(string)
  default     = [ "origins" ]
}
variable "cloudant_cors_allow_credentials" {
  description = "Boolean value to allow authentication credentials. If set to true, browser requests must be done by using withCredentials = true."
  type        = bool
  default     = false
}
variable "cloudant_cors_enable_cors" {
  description = "Boolean value to turn CORS on and off."
  type        = bool
  default     = false
}

// Data source arguments for cloudant_server_information

// Data source arguments for cloudant_activity_tracker_events

// Data source arguments for cloudant_capacity_throughput

// Data source arguments for cloudant_cors
