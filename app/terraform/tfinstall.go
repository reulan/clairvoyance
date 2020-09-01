package terraform

import (
	"io/ioutil"
	"os"

	tfinstall "github.com/hashicorp/terraform-exec/tfinstall"
)

func ConfigureBinary() string {
	tmpDir, err := ioutil.TempDir("", "tfinstall")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tmpDir)

	execPath, err := tfinstall.Find(tfinstall.LatestVersion(tmpDir, false))
	if err != nil {
		panic(err)
	}

	return execPath
}
