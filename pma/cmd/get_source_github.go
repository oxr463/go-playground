package cmd

import (
	"github.com/spf13/cobra"
)

var getSourceGithubCmd = &cobra.Command{
	Use:   "github",
	Short: "Get github source",
	Long:  "",
}

func init() {
	getSourceCmd.AddCommand(getSourceGithubCmd)
}
