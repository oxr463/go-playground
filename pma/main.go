package main

import (
	"github.com/oxr463/pma/cmd"
	_ "github.com/oxr463/pma/pkg/pma"
)

var VERSION = "0.0.1"

func main() {
	cmd.Execute()
}
