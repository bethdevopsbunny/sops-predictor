package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sops-predictor",
	Short: "predict encrypted fields",
	Long:  " ",
}

func init() {

}

func Execute() error {
	return rootCmd.Execute()
}
