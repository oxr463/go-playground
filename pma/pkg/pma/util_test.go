package pma

import (
	"fmt"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log("Testing Package")
	var rke = Package{
		source:  "github",
		url:     "https://github.com/rancher/rke",
		version: "v1.4.1",
	}
	fmt.Println(rke)
	t.Fail()
}

func TestGetGitBranch(t *testing.T) {
	t.Log("Testing Get Git Branch")
	GetGitBranch()
	t.Fail()
}

func TestGetPackageDirectory(t *testing.T) {
	package_directory = GetPackageDirectory("rke")
	if package_directory != "testing/rke" {
		t.Log("Testing Get Package Directory")
		t.Fail()
	}
}
