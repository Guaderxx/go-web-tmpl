package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "0.1.0",
	Long:  `All software has versions. This is Tmpl's`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("Tmpl v0.1.0")
	},
}