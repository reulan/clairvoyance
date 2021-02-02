package terraform

//import "os"

func DetectTerraformBinary() string {
	// Get version of Terraform binary to use (local)
	//var binaryDir = os.Getenv("GOPATH") + "/src/clairvoyance/tfinstall/terraform_" + os.Getenv("CLAIRVOYANCE_TERRAFORM_VERSION")
	//tfBinary := terraform.DetectBinary(binaryDir, terraformVersion)

	// If Linux default to:
	//var tfBinary = "/usr/bin/terraform"

	// If MacOS default to:
	var tfBinary = "/usr/local/bin/terraform"

	return tfBinary

}
