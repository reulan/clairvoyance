package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"clairvoyance/app/extras"
	"clairvoyance/app/general"
	"clairvoyance/app/terraform"
	"clairvoyance/log"
)

func init() {
	log.Debug("[report.go/init] Running CLI command: ")
	rootCmd.AddCommand(reportCmd)
	reportCmd.Flags().StringP("output", "o", "discord", "Choose the target medium to report to. (discord, stdout)")
	//reportCmd.Flags().Bool("festive", true, "Determine if ASCII art + emoji's are printed.")
}

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
	Short: "Reports terraform drift to Discord",
	Long: `Reports terraform drift to Discord
		Usage:
		clairvoyance report`,

	Run: func(cmd *cobra.Command, args []string) {
		optionOutput, _ := cmd.Flags().GetString("output")
		//optionFestive := cmd.Flags().Bool("festive", true, "Prints ASCII art + emojis")

		// Get version of Terraform binary to use
		//var binaryDir = os.Getenv("GOPATH") + "/src/clairvoyance/tfinstall/terraform_" + os.Getenv("CLAIRVOYANCE_TERRAFORM_VERSION")
		//tfBinary := terraform.DetectBinary(binaryDir, terraformVersion)

		//var tfBinary = "/usr/bin/terraform"
		var tfBinary = "/usr/local/bin/terraform"

		// Setup Terraform Version to use
		var _, tfVersionSet = os.LookupEnv("CLAIRVOYANCE_TERRAFORM_VERSION")
		var terraformVersion string
		_ = terraformVersion

		if tfVersionSet {
			terraformVersion = os.Getenv("CLAIRVOYANCE_TERRAFORM_VERSION")
		} else {
			// should be "" or "latest" - will hardcode to latest version for now
			terraformVersion = "0.14.3"
		}

		// Setup projects to plan
		var clarivoyanceProjectDir = os.Getenv("CLAIRVOYANCE_PROJECT_DIR")
		projects, err := general.FindPlannableProjects(clarivoyanceProjectDir, "*.tf")
		if err != nil {
			panic(err)
		}

		// override projects for testing
		projects = []string{"/Users/mpmsimo/noobshack/gameservers/rust/rustdm"}

		/* Terraform Drift Report */
		driftDetectTime := time.Now()
		var terraformServices []*terraform.TerraformService
		tfChannel := make(chan *terraform.TerraformService)

		// Check to see if projects list is more than 0, to determine if plannable
		var PlannableProjects bool

		if len(projects) == 0 {
			PlannableProjects = false
		} else {
			PlannableProjects = true
		}

		if PlannableProjects {
			for _, absProjectPath := range projects {
				go terraform.GetProjectDrift(tfChannel, absProjectPath, tfBinary)
			}

			for _, _ = range projects {
				terraformServices = append(terraformServices, <-tfChannel)
			}
		} else {
			log.Printf("[reportCmd] No *.tf files found in CLAIRVOYANCE_WORKING_DIR.")
		}

		log.Printf("[reportCmd] Drift report took %s to run.\n", time.Since(driftDetectTime))

		// Festive! (for HashiCorp Holiday Hackstravaganza)
		fmt.Println(extras.GetAsciiArt())
		fmt.Println(extras.Snowflakes)

		// Drift Report
		terraform.CreateTableStdout(terraformServices)

		// Where is the message going?
		if optionOutput == "stdout" || optionOutput == "" {
			log.Debug("[cmdReport] Outputting to Stdout.")
		} else if optionOutput == "discord" {
			log.Debug("[cmdReport] Outputting to Discord.")
			//reporting.SendMessageDiscord(message)
		} else {
			log.Errorf("[cmdReport] optionOutput: [%s] not supported (discord, stdout)", optionOutput)
		}
	},
}
