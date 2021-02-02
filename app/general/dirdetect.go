package general

import (
	"os"
	"path/filepath"

	"clairvoyance/log"
)

var prefixPath string = os.Getenv("CLAIRVOYANCE_PROJECT_DIR")

func FindPlannableProjects(baseDir string, pattern string) ([]string, error) {
	var projects []string
	var prefixPath string = os.Getenv("CLAIRVOYANCE_PROJECT_DIR")

	// For the specified baseDir dir, find ALL files within all directories.
	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if file is directory, if so skip
		if info.IsDir() {
			return nil
		}

		// If the directory contains *.tf files, take note of:
		// - Directory containing those said *.tf files (relative path)
		// - If the directory has a file which contains `required_version` scrape dat.
		// TODO: Don't use rel/abs as done with filepath, use GetProjectName function (to get shorthand)
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			relFile, err := filepath.Rel(baseDir, path)
			if err != nil {
				panic(err)
			}
			projectPlanDir := (prefixPath + "/" + filepath.Dir(relFile))

			// check if dir is unique
			if contains(projects, projectPlanDir) {
				log.Printf("[WalkMatch] Project %s exists in []string array.", projectPlanDir)
			} else {
				projects = append(projects, projectPlanDir)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return projects, err
}

// Check to see if projectDir is in projects array
func contains(projects []string, projectDir string) bool {
	for _, dir := range projects {
		if dir == projectDir {
			log.Printf("[contains] %s exists in projects []string.", dir)
			return true
		}
	}
	return false
}
