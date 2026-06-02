include "root" {
  path = find_in_parent_folders("root.hcl")
}

include "defaults" {
  path = find_in_parent_folders("_default_params/service-account.hcl")
}

locals {
  project_id = yamldecode(file(find_in_parent_folders("project.yaml"))).project_id
}

inputs = {
  project_id = local.project_id
}

terraform {
  source = "../modules/service-account"
}
