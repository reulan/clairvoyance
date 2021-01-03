package terraform

/*

import (
	tfinstall "github.com/hashicorp/terraform-exec/tfinstall"
)

func DetectBinary(installDir string, version string) string {

	// If terraform is not detected, install to the installation directory on the user's behalf.
	if version == "" {
		tfbinary, err := tfinstall.Find(tfinstall.LatestVersion(installDir, false))

		if err != nil {
			panic(err)
		}

		return tfbinary

	} else {

		tfbinary, err := tfinstall.Find(tfinstall.ExactVersion(version, installDir))
		if err != nil {
			panic(err)
		}

		return tfbinary

	}

	// return nothing if not set
	return ""
}
*/
