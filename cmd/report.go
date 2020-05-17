package cmd

import (
	"clairvoyance/log"
	"fmt"

	"github.com/spf13/cobra"

	"clairvoyance/app/reporting"
	"clairvoyance/app/terraform"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Reports terraform drift to Discord",
	Long: `Reports terraform drift to Discord
		Usage:
		clairvoyance report`,

	Run: func(cmd *cobra.Command, args []string) {
			// TODO: Implement parsing of atlantis.yaml to know where to search for Terraform projects
			tfexecCfg := terraform.Configure()
			terraform.Init(tfexecCfg)
			tfState := terraform.Show(tfexecCfg)

			outputs := tfState.Values.Outputs
			log.Info("Outputs:\n")
			log.Info(outputs)

			message := fmt.Sprintf("Project: [%s] is running version Terraform %s.", tfexecCfg.WorkingDir, tfState.TerraformVersion)
			log.Info("Message:\n")
			log.Info(message)
			var sendMessage string

			optDebug, _ := cmd.Flags().GetBool("debug")
			if optDebug {
				sendMessage = reporting.DebugFormatMessage()
			} else {
				sendMessage = reporting.FormatDriftReport(message)
			}
			log.Info("SendMessage:\n")
			log.Info(sendMessage)
			reporting.SendReport(sendMessage)
		},
}

func init() {
	fmt.Println("cmd/report/go running.")
	rootCmd.AddCommand(reportCmd)
	reportCmd.Flags().BoolP("debug", "d", false, "Sends a debug message to the channel instead of the drift report.")
}
