// Resource arguments for ibmcloud_api_key
variable "ibmcloud_api_key" {
  description = "IBM Cloud API key."
  type        = string
}

// Resource arguments for service_region
variable "service_region" {
  description = "Region in which service has to be provisioned."
  type        = string
  default     = "us-south"
}

// Resource arguments for cloudant_dedicated_hardware_name
variable "cloudant_dedicated_hardware_name" {
  description = "The actual physical location of the Dedicated Hardware plan instance."
  type = string
}
