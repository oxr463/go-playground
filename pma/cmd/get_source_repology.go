package cmd

import (
	"github.com/spf13/cobra"
)

var getSourceRepologyCmd = &cobra.Command{
	Use:   "repology",
	Short: "Get repology source",
	Long:  "",
}

func init() {
	getSourceCmd.AddCommand(getSourceRepologyCmd)
}
