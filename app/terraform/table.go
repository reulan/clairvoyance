package terraform

import (
	"os"
	"strconv"

	//"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/rodaine/table"
)

// Get Terraform Drift report structure
// Unpack struct keys as header
// Insert values into data 2D string array
// Concatenate all common values together and ring up total

// this is ascii
func CreateTable(ts *TerraformService) {
	data := [][]string{
		[]string{ts.ProjectName, ts.TerraformVersion, strconv.Itoa(ts.CountAdd), strconv.Itoa(ts.CountChange), strconv.Itoa(ts.CountDestroy), ts.Summary},
		[]string{"asdf", "asdf", "asdf", "asdf", "asdf", "placeholder"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Project Name", "Version", "Add", "Change", "Delete", "Information"})
	table.SetFooter([]string{"", "Total", "aggregatedAdd", "aggChange", "aggDelete", ""})
	table.SetBorder(false) // Set Border to false
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}

// this is meant for stdout to allow for easier text manipluation

func CreateTableStdout(ts *TerraformService) {
	//headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	//columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Project Name", "Version", "Add", "Change", "Delete", "Information")
	//tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	tbl.AddRow(ts.ProjectName, ts.TerraformVersion, strconv.Itoa(ts.CountAdd), strconv.Itoa(ts.CountChange), strconv.Itoa(ts.CountDestroy), ts.Summary)
	tbl.Print()
}
