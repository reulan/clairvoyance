package terraform

import (
//tfinstall "github.com/hashicorp/terraform-exec/tfinstall"
//"clairvoyance/log"
)

/*
// If terraform is not detected, install to the installation directory on the user's behalf.
func DetectBinary(installDir string, version string) string {
	if version == "" {
		tfbinary, err := tfinstall.Find(tfinstall.LatestVersion(installDir, false))
		if err != nil {
			log.Errorf("[DetectBinary] Could not install %s to %s.", version, installDir)
		}
		return tfbinary
	} else {
		tfbinary, err := tfinstall.Find(tfinstall.ExactVersion(version, installDir))
		if err != nil {
			log.Errorf("[DectectBinary] Could not install %s to %s.", version, installDir)
		}
		return tfbinary
	}
	return ""
}
*/
