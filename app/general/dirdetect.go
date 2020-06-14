package general

import (
	"os"

	"clairvoyance/app/terraform"
	"clairvoyance/log"
)

//TODO: Fix all dirdetect functionalities, this is just a basic first pass on what I would like to support.

// Single directory
func DetectService(path string) string {
	var project string

	includesHCL := terraform.IdentifyHCL(path)
	if includesHCL {
		project = path
	}

	return project
}

// Recursive directory search - not working, needs big update
func DetectServices(path string) []string {
	var services []string

	dir, err := os.Open(path)
	if err != nil {
		log.Errorf("Failed to open directory: %s", dir)
	}


	includesHCL := terraform.IdentifyHCL(path)
	if includesHCL {
		//services = path
	}

	log.Infof("%s", services)

	return services
}

// TODO: Terraform Cloud Workspace detection
