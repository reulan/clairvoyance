# clairvoyance
Drift detection and reporting for Terraform.

## Overview
Currently, the `clairvoyance` software does the following:
- Identifies installed version of Terraform or installs specified version.
- Configure Terraform Project working dir
- Initialize and show the statefile information for the Terraform project
- Reporting to Discord text channel via webhook or standard output

In the future I would like to support:
- Planning multiple states across backends
- Terraform project detection (local file, atlantis.yaml, Terraform Cloud workspaces)
- Terraform Report stats (added/changed/deleted, total projects, versions, etc)
- Clarivoyance metadata (how long it takes for a plan or report to be completed + app metrics)
- Generate HCL code suggestions
- Report to other mediums (Slack, IRC, email)

## Project setup
Ensure Golang is installed and configured.

To reinitialize the modules and recreate the dependency tree the following can be done:
`make deps`

### Setting Environment variables
The following environment variables will need to be set for `clairvoyance` to run:
- `CLAIRVOYANCE_ATLANTIS_YAML` (path to `atlantis.yaml`)
- `CLAIRVOYANCE_TERRAFORM_VERSION` (version of terraform binary to use by default, if not specified)
- `CLAIRVOYANCE_WORKING_DIR` (path to terraform service to plan) 
- `DISCORD_WEBHOOK_CHANNEL` (discord channel name. e.x. `#clairvoyance`)
- `DISCORD_WEBHOOK_SECRET`

The Discord secret expects to contain everything after the webhooks route:
`https://discordapp.com/api/webhooks/$DISCORD_WEBHOOK_SECRET`

### Build and Run
#### tfinstall
In order for `clairvoyance` to run, a path to a Terraform binary must be specified.
Currently the application will look in it's own directory at the `./tfinstall` directory.
```
make tfinstall
```

When a report is generated via `make`, the binary will be build before executed.
```
make report-stdout
OR
make report-discord
`````

### Update Version
Modify `version/version.go` and add a major, minor or patch version based off contributions.

## Additional information
This repository was bootstrapped with [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang).

### Notable packages
Packages can be downloaded from public GitHub repositories, like so:
`go get https://github.com/$USER/$REPO`

Modules that are intended to be used are documented below.
- [hclwrite](https://github.com/hashicorp/hcl/tree/v2.0.0/hclwrite) - write HCL on the fly
- [terrafmt](https://github.com/terrycain/terrafmt) - format the HCL output, if live update is used
- [terraform-exec](https://github.com/kmoe/terraform-exec) - so we can init/plan/apply via the Terraform CLI programmatically.
- [tfvar](https://github.com/shihanng/tfvar) - programatic definition and generation of variables based on user input
