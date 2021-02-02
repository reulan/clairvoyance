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
	var fst int = 0
	var nct int = 0

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	driftDetectedTable := table.New("Project Name", "Version", "Add", "Change", "Delete", "Information")
	driftDetectedTable.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	failedServicesTable := table.New("Project Name", "Version", "Information")
	failedServicesTable.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	noChangesTable := table.New("Project Name", "Version", "Information")
	noChangesTable.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for _, service := range tsArray {
		if service.Summary == "No changes." {
			noChangesTable.AddRow(service.ProjectName, service.TerraformVersion, service.Summary)
			log.Debug("[CreateTableStdout] Terraform service contains no drift.")
			nct++
		} else if service.Summary == "Drift detected for Plan." {
			driftDetectedTable.AddRow(service.ProjectName, service.TerraformVersion, strconv.Itoa(service.CountAdd), strconv.Itoa(service.CountChange), strconv.Itoa(service.CountDestroy), service.Summary)
			log.Debugf("[CreateTableStdout] Added %s to driftDetectedTable.", service.ProjectName)
			ddt++
		} else {
			failedServicesTable.AddRow(service.ProjectName, service.TerraformVersion, strconv.Itoa(service.CountAdd), strconv.Itoa(service.CountChange), strconv.Itoa(service.CountDestroy), service.Summary)
			log.Debugf("[CreateTableStdout] Added %s to failedServicesTable.", service.ProjectName)
			fst++
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
	if fst >= 1 {
		failedServicesTable.Print()
		fmt.Println("")
	}
	log.Debug("Sent Drift Report tables to stdout.")
}

func FailedServicesTable(failedServices []string) {
	headerFmt := color.New(color.FgRed, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgMagenta).SprintfFunc()

	failedServicesTable := table.New("Project Name")
	failedServicesTable.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, service := range failedServices {
		failedServicesTable.AddRow(service)
		log.Debugf("[FailedServices] Added %s to failedServicesTable.", service)
	}

	// Find a better way omit tables
	fmt.Println("")
	failedServicesTable.Print()
	fmt.Println("")
	log.Debug("Sent Failed Services Table to stdout.")
}
