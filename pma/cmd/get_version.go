package cmd

import (
	"github.com/spf13/cobra"
)

var getVersionCmd = &cobra.Command{
	Use:     "versions",
	Aliases: []string{"version"},
	Short:   "Get package versions",
	Long:    "",
}

func init() {
	getCmd.AddCommand(getVersionCmd)
}
