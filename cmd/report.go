package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"clairvoyance/app/extras"
	"clairvoyance/app/general"
	"clairvoyance/app/terraform"
	"clairvoyance/log"
)

/*
In order for a report to be done, a tfexec config should be populated and we need to ensure the following
values have been captured.

The following options for additional reporting functionality.
	clairvoyance report:
		--command <show/plan/apply> (Performs limited Terraform CLI logic, a more comprehensive report behaviour is used)
		--path <working_directory>
		--output [<discord>, <stdout>]
		--festive

		TODO: *what does a config file look like, where is this loaded from? (based off tfexc cfg?)
		--config <clairvoyance_config>

	clairvyoance report --output discord --festive
	clairvoyance report --path ~/noobshack --output discord
	clairvoyance report --command show --path ~/noobshack --output stdout
*/

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Outputs Terraform Drift Report to specified medium",
	Long: `Outputs Terraform Drift Report to specified medium"
		Usage:
		clairvoyance report`,

	Run: func(cmd *cobra.Command, args []string) {
		// Setup CLI options
		optionOutput, _ := cmd.Flags().GetString("output")
		optionFestive, _ := cmd.Flags().GetBool("festive")

		// Configure Terraform settings for Clairvoyance
		tfVersion := terraform.DetectTerraformVersion()

		tfBinary := terraform.DetectBinary(tfVersion)
		cvProjects, cvIsPlannable := general.GetPlannableProjects()

		// Terraform Drift Report
		driftDetectTime := time.Now()

		var terraformServices []*terraform.TerraformService
		tfChannel := make(chan *terraform.TerraformService)

		if cvIsPlannable {
			for _, absProjectPath := range cvProjects {
				go terraform.GetProjectDrift(tfChannel, absProjectPath, tfBinary)
			}

			for _, _ = range cvProjects {
				terraformServices = append(terraformServices, <-tfChannel)
			}
		} else {
			log.Printf("[reportCmd] No *.tf files found in CLAIRVOYANCE_WORKING_DIR.")
		}

		// Where is the message going?
		if optionOutput == "stdout" || optionOutput == "" {
			log.Debug("[cmdReport] Outputting to Stdout.")

			// Festive! (for HashiCorp Holiday Hackstravaganza)
			if optionFestive {
				fmt.Println(extras.GetAsciiArt())
				fmt.Println(extras.Snowflakes)
			}
			terraform.CreateTableStdout(terraformServices)

			// Drift Report
			log.Printf("[reportCmd] Drift report took %s to report to stdout.\n", time.Since(driftDetectTime))
		} else if optionOutput == "discord" {
			log.Debug("[cmdReport] Outputting to Discord.")
			//reporting.SendMessageDiscord(message)
		} else {
			log.Errorf("[cmdReport] optionOutput: [%s] not supported (discord, stdout)", optionOutput)
		}
	},
}

func init() {
	log.Debug("[report.go/init] Running CLI command: ")
	rootCmd.AddCommand(reportCmd)
	reportCmd.Flags().StringP("output", "o", "stdout", "Where to report to. (stdout, discord)")
	reportCmd.Flags().BoolP("festive", "f", false, "Determine if ASCII art + emoji's are printed.")
}
