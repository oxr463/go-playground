package cmd

import (
	"github.com/spf13/cobra"
)

var getSourceCmd = &cobra.Command{
	Use:     "sources",
	Aliases: []string{"source"},
	Short:   "Get package sources",
	Long:    "",
}

func init() {
	getCmd.AddCommand(getSourceCmd)
}
