package pma

import (
	"os"
)

type package struct {
	source  string
	url     string
	version string
}

func GetGitBranch() (git_branch string) {
	// Default branch name
	git_branch = "master"

	// Attempt to set branch from environment variable
	// See: https://git-scm.com/docs/git-var#_variables

	GIT_DEFAULT_BRANCH := os.Getenv("GIT_DEFAULT_BRANCH")

	if GIT_DEFAULT_BRANCH != "" {
		git_branch = GIT_DEFAULT_BRANCH
	}

	return git_branch
}

func GetPackageDirectory(maintainer_package string) (package_directory string) {
	_ = maintainer_package
	package_directory = "FIXME"
	return package_directory
}
