# terragrunt SA
resource "google_service_account" "lucky_function_terragrunt" {
  account_id   = "lucky-function-terragrunt-sa"
  display_name = "Lucky Function service account"
}

resource "google_project_iam_member" "lucky_function_viewer" {
  project = var.project_id
  role    = "roles/viewer"
  member  = "serviceAccount:${google_service_account.lucky_function_terragrunt.email}"
}

# demo resources
# resource "google_service_account" "lucky_function" {
#   account_id   = "lucky-function-sa"
#   display_name = "Lucky Function service account"
# }

# resource "google_project_iam_member" "lucky_function_storage_viewer" {
#   project = var.project_id
#   role    = "roles/storage.objectViewer"
#   member  = "serviceAccount:${google_service_account.lucky_function.email}"
# }

# resource "google_project_iam_member" "lucky_function_log_writer" {
#   project = var.project_id
#   role    = "roles/logging.logWriter"
#   member  = "serviceAccount:${google_service_account.lucky_function.email}"
# }
