package cmd

import (
	"github.com/spf13/cobra"
)

var bumpCmd = &cobra.Command{
	Use:   "bump",
	Short: "",
	Long:  "",
}

func init() {
	rootCmd.AddCommand(bumpCmd)
}
