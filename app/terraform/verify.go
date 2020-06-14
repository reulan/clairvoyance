package terraform

import "clairvoyance/log"

func IdentifyHCL(path string) bool {
	log.Info("IdentifyHCL - Ensuring service contains .tf or .tfvars files.")
	var isPlannable bool
	// Identify if a HCL files ending in .tf or .tfvars and is plannable by Terraform
	isPlannable = true

	return isPlannable
}
