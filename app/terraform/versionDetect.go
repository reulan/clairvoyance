package terraform

import "os"

func DetectTerraformVersion() string {
	var terraformVersion string = "0.14.5"

	// Setup Terraform Version to use
	var _, tfVersionSet = os.LookupEnv("CLAIRVOYANCE_TERRAFORM_VERSION")

	if tfVersionSet {
		terraformVersion = os.Getenv("CLAIRVOYANCE_TERRAFORM_VERSION")
	}

	return terraformVersion
}
