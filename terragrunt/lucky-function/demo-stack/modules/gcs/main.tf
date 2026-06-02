terraform {
  backend "gcs" {}
}

provider "google" {
  project = var.project_id
  region  = var.region
}

resource "google_storage_bucket" "demo" {
  name          = var.bucket_name
  project       = var.project_id
  location      = var.location
  force_destroy = false
}

resource "google_storage_bucket_iam_member" "sa_viewer" {
  bucket = google_storage_bucket.demo.name
  role   = "roles/storage.objectViewer"
  member = "serviceAccount:${var.service_account_email}"
}

output "bucket_name" {
  value = google_storage_bucket.demo.name
}
