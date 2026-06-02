terraform {
  backend "gcs" {}
}

provider "google" {
  project = var.project_id
  region  = var.region
}

resource "google_service_account" "demo" {
  project      = var.project_id
  account_id   = var.service_account_id
  display_name = var.display_name
}

output "email" {
  value = google_service_account.demo.email
}
