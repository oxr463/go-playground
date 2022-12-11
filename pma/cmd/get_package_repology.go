package cmd

import (
	"github.com/spf13/cobra"
)

var getPackageRepologyCmd = &cobra.Command{
	Use:   "repology",
	Short: "Get maintainer repology packages",
	Long:  "",
}

func init() {
	getPackageCmd.AddCommand(getPackageRepologyCmd)
}
