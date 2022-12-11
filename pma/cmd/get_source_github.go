package cmd

import (
	"github.com/spf13/cobra"
)

var getSourceGithubCmd = &cobra.Command{
	Use:   "github",
	Short: "Get source github",
	Long:  "",
}

func init() {
	getSourceCmd.AddCommand(getSourceGithubCmd)
}
