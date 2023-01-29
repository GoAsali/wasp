package main

import (
	"main/lib/cli"
	"os"
)

func main() {
	args := os.Args[1:]

	commandCli := cli.NewCommandCli()
	if len(args) == 0 {
		commandCli.PrintHelp()
		return
	}
}
