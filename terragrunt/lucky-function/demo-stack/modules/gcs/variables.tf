variable "project_id" {
  description = "GCP project ID"
  type        = string
}

variable "region" {
  description = "GCP region"
  type        = string
}

variable "bucket_name" {
  description = "GCS bucket name"
  type        = string
}

variable "location" {
  description = "GCS bucket location"
  type        = string
}

variable "service_account_email" {
  description = "Service account email for bucket IAM"
  type        = string
}
