package main

import "syspeak/cmd"

func main() {
	if err := cmd.SysPeakCmd.Execute(); err != nil {
		panic(err)
	}
}
