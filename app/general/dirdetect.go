package general

import (
	//"log"
	"os"
	"path/filepath"
)

// Check out this way to use a map[string]bool
// https://play.golang.org/p/qw2FG5a9hv_Q

func FindPlannableProjects(root string, pattern string) ([]string, error) {
	var projects []string

	// For the specified root dir, find ALL files within all directories.
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check
		if info.IsDir() {
			return nil
		}

		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			relFile, err := filepath.Rel(root, path)
			if err != nil {
				panic(err)
			}
			projectPlanDir := filepath.Dir(relFile)

			// check if dir is unique
			if contains(projects, projectPlanDir) {
				//log.Printf("[WalkMatch] Project %s exists in []string array.", projectPlanDir)
			} else {
				projects = append(projects, projectPlanDir)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return projects, nil
}

// Check to see if projectDir is in projects array
func contains(projects []string, projectDir string) bool {
	for _, dir := range projects {
		if dir == projectDir {
			return true
		}
	}
	return false
}
