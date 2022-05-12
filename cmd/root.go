package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var list = []string{}

var rootCmd = &cobra.Command{
	Use:   "rob-clip",
	Short: "Rob-clip is tool that will help you with your clipboard history",
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(stopCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}