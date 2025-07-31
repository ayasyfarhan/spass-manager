package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spass-manager",
	Short: "A program to decrypt .spass files",
	Long:  `A program to decrypt .spass files and export them to different formats`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
