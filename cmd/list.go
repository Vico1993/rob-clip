package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "init",
	Short: "Runs initialization of Jobs",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Je list des truc un jours")
	},
}