package cmd

import (
	"clairvoyance/log"
	"fmt"

	"github.com/spf13/cobra"

	"clairvoyance/app/reporting"
	"clairvoyance/app/terraform"
)

/*
In order for a report to be done, a tfexec config should be populated and we need to ensure the following
values have been captured.

The following options for additional reporting functionality.
	clairvoyance report:
		--command <init/plan/apply/show> (Performs limited Terraform CLI logic, a more comprehensive report behaviour is used)
		--path <working_directory>
		--output [<discord>, <stdout>]
		--debug (overrides --output)

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
		//optOutput, _ := cmd.Flags().GetString("output")
		optDebug, _ := cmd.Flags().GetBool("debug")

		tfexecCfg := terraform.ConfigureTerraform()
			terraform.Init(tfexecCfg)

			tfState := terraform.Show(tfexecCfg)

			outputs := tfState.Values.Outputs
			log.Info("Outputs:\n")
			log.Info(outputs)
		///Users/mpmsimo/noobshack/valorant
			message := fmt.Sprintf("Project: [%s] is running version Terraform %s.", tfexecCfg, tfState.TerraformVersion)
			log.Info("Message:\n")
			log.Info(message)
			var sendMessage string


		if optDebug {
				sendMessage = reporting.DebugFormatMessage()
			} else {
				sendMessage = reporting.FormatDriftReport(message)
			}
			log.Info("SendMessage:\n")
			log.Info(sendMessage)
			reporting.SendMessageDiscord(sendMessage)
		},
}

func init() {
	fmt.Println("cmd/report/go running.")
	rootCmd.AddCommand(reportCmd)
	reportCmd.Flags().BoolP("debug", "d", false, "Sends a debug message to the channel instead of the drift report.")
}
