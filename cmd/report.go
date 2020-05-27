package cmd

import (
	"clairvoyance/log"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"clairvoyance/app/reporting"
	"clairvoyance/app/terraform"
)

/*
In order for a report to be done, a tfexec config should be populated and we need to ensure the following
values have been captured.

The following options for additional reporting functionality.
	clairvoyance report:
		--command <show/plan/apply> (Performs limited Terraform CLI logic, a more comprehensive report behaviour is used)
		--path <working_directory>
		--output [<discord>, <stdout>]

		TODO: *what does a config file look like, where is this loaded from? (based off tfexc cfg?)
		--config <clairvoyance_config>

	clairvoyance report --path ~/noobshack --output discord
	clairvoyance report --command show --path ~/noobshack --output stdout
 */

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Reports terraform drift to Discord",
	Long: `Reports terraform drift to Discord
		Usage:
		clairvoyance report`,

	Run: func(cmd *cobra.Command, args []string) {
		//optCommand, _ := cmd.Flags().GetString("command")
		//optPath, _ := cmd.Flags().GetString("path")
		optOutput, _ := cmd.Flags().GetString("output")

		// configure service - referring to a tfexec config (a single terraform project definition)
		//TODO: copy files over to the container
		var workingDir = os.Getenv("GOPATH") + "/src/github.com/kmoe/terraform-exec/testdata"
		optPath, _ := filepath.Abs(workingDir)
		service := terraform.ConfigureTerraform(optPath)
		terraform.Init(service)

		// Bypass optPath as only Show is supported as of now
		state := terraform.Show(service)
		formattedOutput := reporting.FormatTerraformShow(state)

		// Where is the message going?
		if optOutput == "discord" {
			//reporting.SendMessageDiscord(formattedOutput)
			} else if optOutput == "stdout" {
				//reporting.SendMessageStdout(formattedOutput)
				reporting.SendJSONStdout(formattedOutput)
		} else {
				log.Errorf("cmd/report - optOutput: [%s] not supported (discord, stdout)", optOutput)
			}
		},
}

func init() {
	fmt.Println("cmd/report/go running.")
	rootCmd.AddCommand(reportCmd)
	//reportCmd.Flags().StringP("command", "c", "show", "Performs a specific Terraform command against the given project. (defaults to Show)")
	//reportCmd.Flags().StringP("path", "p", "/path/to/terraform/project", "Specify the path of the Terraform project you'd like to report on")
	reportCmd.Flags().StringP("output", "o", "discord", "Choose the target medium to report to. (discord, stdout)")
}
