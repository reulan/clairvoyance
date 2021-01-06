package terraform

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

// this is meant for stdout to allow for easier text manipluation
func CreateTableStdout(tsArray []*TerraformService) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Project Name", "Version", "Add", "Change", "Delete", "Information")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for _, service := range tsArray {
		tbl.AddRow(service.ProjectName, service.TerraformVersion, strconv.Itoa(service.CountAdd), strconv.Itoa(service.CountChange), strconv.Itoa(service.CountDestroy), service.Summary)
	}

	fmt.Println("")
	tbl.Print()
}
