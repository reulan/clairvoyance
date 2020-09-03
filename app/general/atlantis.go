package general

// import "gopkg.in/yaml.v2"

// https://github.com/runatlantis/atlantis/tree/master/server/events/yaml

// Parses atlantis.yaml to figure out which Terraform service to use.
// By default loads tftest/atlantis.yaml or CLAIRVOYANCE_ATLANTIS_FILE if specified.
// search in projects, get dir/terraform_version
// return service_name, terraform_version
