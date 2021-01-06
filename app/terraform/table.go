package terraform

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/rodaine/table"

	"clairvoyance/log"
)

// this is meant for stdout to allow for easier text manipluation
func CreateTableStdout(tsArray []*TerraformService) {
	var ddt int = 0
	var nct int = 0

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	driftDetectedTable := table.New("Project Name", "Version", "Add", "Change", "Delete", "Information")
	driftDetectedTable.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	noChangesTable := table.New("Project Name", "Version", "Information")
	noChangesTable.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for _, service := range tsArray {
		if service.Summary == "No changes." {
			noChangesTable.AddRow(service.ProjectName, service.TerraformVersion, service.Summary)
			log.Debug("[CreateTableStdout] Terraform service contains no drift.")
			nct++
		} else {
			driftDetectedTable.AddRow(service.ProjectName, service.TerraformVersion, strconv.Itoa(service.CountAdd), strconv.Itoa(service.CountChange), strconv.Itoa(service.CountDestroy), service.Summary)
			log.Debug("[CreateTableStdout] Added %s to driftDetectedTable.", service.ProjectName)
			ddt++
		}

	}

	// Find a better way omit tables
	fmt.Println("")
	if ddt >= 1 {
		driftDetectedTable.Print()
		fmt.Println("")
	}
	if nct >= 1 {
		noChangesTable.Print()
		fmt.Println("")
	}
	log.Debug("Sent Drift Report tables to stdout.")
}
