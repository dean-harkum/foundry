include "root" {
  path = find_in_parent_folders("root.hcl")
}

include "defaults" {
  path = find_in_parent_folders("_default_params/gcs.hcl")
}

locals {
  project_id = yamldecode(file(find_in_parent_folders("project.yaml"))).project_id
}

dependency "service_account" {
  config_path = "../service-account"
  mock_outputs = {
    email = "demo-stack-sa@${local.project_id}.iam.gserviceaccount.com"
  }
}

inputs = {
  project_id           = local.project_id
  service_account_email = dependency.service_account.outputs.email
}

terraform {
  source = "../modules/gcs"
}
