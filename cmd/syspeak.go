package cmd

import "github.com/spf13/cobra"

var SysPeakCmd = &cobra.Command{
	Use:   "syspeak",
	Short: "Syspeak is a runtime service that monitors the performance of the system.",
}

func init() {
	SysPeakCmd.AddCommand(ServerCmd)
}
