package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"clairvoyance/app/reporting"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Reports terraform drift to Discord",
	Long: `Usage:
            clairvoyance report

		   Sends a drift report to Discord.`,
	Run: func(cmd *cobra.Command, args []string) {
		reporting.SendReport()
	},
}

func init() {
	fmt.Println("cmd/report/go running.")
	rootCmd.AddCommand(reportCmd)
}
