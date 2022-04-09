package cmd

import (
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop deamon",
	Run: func(cmd *cobra.Command, args []string) {

	},
}