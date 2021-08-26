// This allows cloudant data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "" {
  value       = ibm_cloudant.cloudant
  description = "cloudant resource instance"
}
