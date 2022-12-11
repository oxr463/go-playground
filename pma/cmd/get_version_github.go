package cmd

import (
	"github.com/spf13/cobra"
)

var getVersionGithubCmd = &cobra.Command{
	Use:   "github",
	Short: "Get github package versions",
	Long:  "",
}

func init() {
	getVersionCmd.AddCommand(getVersionGithubCmd)
}
