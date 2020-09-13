package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"clairvoyance/log"
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

		var workingDir = os.Getenv("CLAIRVOYANCE_WORKING_DIR")
		var binaryDir = os.Getenv("GOPATH") + "/src/clairvoyance/tfinstall/terraform_0.13.2"
		var _, tfVersionSet = os.LookupEnv("CLAIRVOYANCE_TERRAFORM_VERSION")
		var terraformVersion string

		if tfVersionSet {
			terraformVersion = os.Getenv("CLAIRVOYANCE_TERRAFORM_VERSION")
		} else {
			// should be "" or "latest" - will hardcode to latest version for now
			terraformVersion = "0.13.2"
		}

		execPath := terraform.DetectBinary(binaryDir, terraformVersion)
		optPath, _ := filepath.Abs(workingDir)

		// Invoke Terraform
		service := terraform.ConfigureTerraform(optPath, execPath)
		terraform.Init(service)
		state := terraform.Show(service)
		log.Printf("[Show]\n%s", state)
		hasChanges := terraform.Plan(service)
		log.Printf("[Plan]\n%s", hasChanges)

		// Format Terraform output
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
