include {
  path = find_in_parent_folders("root.hcl")
}

locals {
  project_id = yamldecode(file("${get_terragrunt_dir()}/project.yaml")).project_id
}

inputs = {
  project_id = local.project_id
}
