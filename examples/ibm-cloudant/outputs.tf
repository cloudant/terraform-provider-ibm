// This allows cloudant_activity_tracker_events data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_cloudant_activity_tracker_events" {
  value       = ibm_cloudant_activity_tracker_events.cloudant_activity_tracker_events_instance
  description = "cloudant_activity_tracker_events resource instance"
}
// This allows cloudant_capacity_throughput data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_cloudant_capacity_throughput" {
  value       = ibm_cloudant_capacity_throughput.cloudant_capacity_throughput_instance
  description = "cloudant_capacity_throughput resource instance"
}
// This allows cloudant_cors data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_cloudant_cors" {
  value       = ibm_cloudant_cors.cloudant_cors_instance
  description = "cloudant_cors resource instance"
}
