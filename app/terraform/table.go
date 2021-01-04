package terraform

import (
	"strconv"

	//"github.com/fatih/color"
	//"github.com/olekukonko/tablewriter"
	"github.com/rodaine/table"
)

// Get Terraform Drift report structure
// Unpack struct keys as header
// Insert values into data 2D string array
// Concatenate all common values together and ring up total

/*
// this is ascii
func CreateTable(tsArray []*TerraformService) {

	// for each projects add a data row
	data := [][]string{
		[]string{ts.ProjectName, ts.TerraformVersion, strconv.Itoa(ts.CountAdd), strconv.Itoa(ts.CountChange), strconv.Itoa(ts.CountDestroy), ts.Summary},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Project Name", "Version", "Add", "Change", "Delete", "Information"})
	table.SetFooter([]string{"", "Total", string(totalToAdd), string(totalToChange), string(totalToDelete), ""})
	table.SetBorder(false) // Set Border to false
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}
*/

// this is meant for stdout to allow for easier text manipluation
func CreateTableStdout(tsArray []*TerraformService) {
	tbl := table.New("Project Name", "Version", "Add", "Change", "Delete", "Information")

	for _, service := range tsArray {
		tbl.AddRow(service.ProjectName, service.TerraformVersion, strconv.Itoa(service.CountAdd), strconv.Itoa(service.CountChange), strconv.Itoa(service.CountDestroy), service.Summary)
	}

	tbl.Print()
}
