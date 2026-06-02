remote_state {
  backend = "gcs"
  config = {
    project     = "${get_env("GOOGLE_PROJECT_VAR")}"
    bucket      = "${get_env("GOOGLE_PROJECT_VAR")}-tfstate"
    location    = "us-west1"
    prefix      = "${basename(get_terragrunt_dir())}/terraform.tfstate"
    credentials = "${get_env("GOOGLE_APPLICATION_CREDENTIALS") }"
  }
}

terraform {
  extra_arguments "default" {
    commands = ["plan", "apply"]
  }
}

inputs = {
  region = "us-west1"
}
