package cmd

import (
	"github.com/spf13/cobra"
)

var getPackageCmd = &cobra.Command{
	Use:     "packages",
	Aliases: []string{"package"},
	Short:   "Get maintainer packages",
	Long:    "",
}

func init() {
	getCmd.AddCommand(getPackageCmd)
}
