package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ivanpasq/ccwc/cmd/ccwc/count"
	"github.com/ivanpasq/ccwc/cmd/ccwc/help"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "help" {
		help.Run()
	}

	cFlag := flag.Bool("c", false, "outputs the number of bytes in a file")

	flag.Parse()

	output := ""

	if *cFlag {
		result, err := count.CRun(flag.Args())
		if err != nil {
			println(fmt.Sprintln("ccwc error - %w", err))
			os.Exit(1)
		}

		output = fmt.Sprintln("    ", result, flag.Args()[0])
	}

	println(output)

}
