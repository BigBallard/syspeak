package cmd

import "github.com/spf13/cobra"

var ServerCmd = &cobra.Command{
	Use: "server",
}

func init() {
	ServerCmd.Run = func(cmd *cobra.Command, args []string) {
		println("server called")
	}
}
