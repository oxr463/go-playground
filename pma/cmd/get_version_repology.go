package cmd

import (
	"github.com/spf13/cobra"
)

var getVersionRepologyCmd = &cobra.Command{
	Use:   "repology",
	Short: "Get repology package versions",
	Long:  "",
}

func init() {
	getVersionCmd.AddCommand(getVersionRepologyCmd)
}
