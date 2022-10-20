package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {

	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of sops-predictor",
	Long:  `Print the version number of sops-predictor`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sops-predictor - v0.1")
	},
}
