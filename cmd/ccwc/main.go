package main

import (
	"os"

	"github.com/ivanpasq/ccwc/cmd/ccwc/help"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "help" {
		help.Run()
	}
}
