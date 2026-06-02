output "service_account_email" {
  description = "Email of the terragrunt service account used by lucky-function"
  value       = google_service_account.lucky_function_terragrunt.email
}
