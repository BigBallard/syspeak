package cmd

import (
	"github.com/spf13/cobra"
	"syspeak/internal/server"
)

var ServerCmd = &cobra.Command{
	Use: "server",
}

func init() {
	ServerCmd.Run = func(cmd *cobra.Command, args []string) {
		s := server.NewServerWithOptions()
		s.Start()
		println("Ending command")
	}
}
