# clairvoyance
Drift detection and reporting for Terraform.

Currently, the software does the following:
- Configure Terraform Project with a tf binary + working dir
- Initialize and show the statefile information for the Terraform project
- Reporting to Discord text channel via webhook or standard output

In the future I would like to support:
- Planning multiple states across backends
- Terraform project detection (local file, atlantis.yaml, Terraform Cloud workspaces)
- Terraform Report stats (added/changed/deleted, total projects, versions, etc)
- Clarivoyance metadata (how long it takes for a plan or report to be completed + app metrics)
- Generate HCL code suggestions

## Usage
### Setting Environment variables
The following environment variables will need to be set for `clairvoyance` to run:
- `DISCORD_WEBHOOK_NAME` (just a string, typically the discord channel name. e.x. `#clairvoyance`)
- `DISCORD_WEBHOOK_SECRET`

The Discord secret expects to contain everything after the webhooks route:
`https://discordapp.com/api/webhooks/$DISCORD_WEBHOOK_SECRET`

### Running
See Development/Build and Run below

## Development
This project requires Go to be installed. 
On OS X with Homebrew you can just run `brew install go`.

### Modules
To reinitialize the modules and recreate the dependency tree the following can be done:
```
cd $GOPATH/src/clairvoyance
rm go.mod go.sum
go mod init
go mod tidy
```

### Testing
From the root directory you can run:
`make test`

Currently there are no tests...
However, test driven development should is the way to go forward for this project.

### Build and Run
Run the binary after it's been packaged:
```console
$ make build
$ ./bin/clairvoyance
$ ./bin/clairvoyance
```

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
