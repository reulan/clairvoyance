# clairvoyance
Drift detection and reporting for Terraform.

## Overview
Currently, the `clairvoyance` software does the following:
- Allows running of Terraform commands for specified working directory.
- Show the statefile information and if any changes are detected. 
- Reporting to Discord text channel via webhook or standard output

In the future I would like to support:
- Identifies installed version of Terraform or installs specified version.
- Planning multiple states across backends
- Terraform project detection (local file, atlantis.yaml, Terraform Cloud workspaces)
- Terraform Report stats (added/changed/deleted, total projects, versions, etc)
- Clarivoyance metadata (how long it takes for a plan or report to be completed + app metrics)
- Generate HCL code suggestions
- Report to other mediums (Slack, IRC, email)

## Project setup
Ensure Golang is installed and configured.

### Setting Environment variables
The following environment variables will need to be set for `clairvoyance` to run:
- `CLAIRVOYANCE_TERRAFORM_VERSION` (version of Terraform to use)
- `CLAIRVOYANCE_WORKING_DIR` (path to terraform service to plan) 
- `DISCORD_WEBHOOK_CHANNEL` (just a string, typically the discord channel name. e.x. `#clairvoyance`)
- `DISCORD_WEBHOOK_SECRET`

The Discord secret expects to contain everything after the webhooks route:
`https://discordapp.com/api/webhooks/$DISCORD_WEBHOOK_SECRET`

### Locally building an image
To reinitialize the modules and recreate the dependency tree the following can be done:
- `make deps`
- `make check-env-vars` (in order to see if you have properly configured your system to work with clairvoyance.)

### Setup Terraform Version via Binary (optional)
In order for `clairvoyance` to run, a path to a Terraform binary must be specified.
Currently the application will look in it's own directory at the `./tfinstall` directory.
```
make tfinstall
```

This will dowload the version assocaited with `CLAIRVOYANCE_TERRAFORM_VERSION`.

If `make tfinstall` is not used, then Clarivoyance will fall back to the binary specified on `/usr/bin/terraform`.

### Build and Run
When a report is generated via `make`, the binary will be build before executed.
```
make report-stdout
OR
make report-discord
```

These commands are wrappers around the CLI tool, examples:
- `clairvoyance report stdout`
- `clairvoyance report discord --festive`


## Development
### Update Version
Modify `version/version.go` and add a major, minor or patch version based off contributions.

### Notable packages
Packages can be downloaded from public GitHub repositories, like so:
`go get https://github.com/$USER/$REPO`

Modules that are intended to be used are documented below.
- [hclwrite](https://github.com/hashicorp/hcl/tree/v2.0.0/hclwrite) - write HCL on the fly
- [terrafmt](https://github.com/terrycain/terrafmt) - format the HCL output, if live update is used
- [terraform-exec](https://github.com/kmoe/terraform-exec) - so we can init/plan/apply via the Terraform CLI programmatically.
- [tfvar](https://github.com/shihanng/tfvar) - programatic definition and generation of variables based on user input

## Additional information
This repository was bootstrapped with [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang).
