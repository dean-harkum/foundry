# lucky-function project

Self contained directory Terragrunt directory

## required

The `GOOGLE_PROJECT_VAR` env var will be required for the remote state

## usage

`main` - can run a simple `terragrunt plan` in the directory, which runs Terragrunt for all resources within
`demo-stack` - can run a `terragrunt run-all plan` in the directory, which runs Terragrunt against a stack.
