package cli

import (
	"fmt"
	"github.com/fatih/color"
	"main/lib"
)

type Command struct {
	Short   string
	Command string
	Help    string
}

type CommandCli struct {
	commands []Command
	config   lib.Config
}

func NewCommandCli() CommandCli {
	commands := []Command{
		{Help: "Display help for the given command. When no command is given display help for the list command", Short: "h", Command: "help"},
	}

	return CommandCli{commands: commands, config: lib.GetConfig()}
}

func (c *CommandCli) PrintHelp() {

	fmt.Print("Go Asali Framework ")
	color.Green("%s\n", c.config.Version)
	fmt.Println()
	color.Yellow("Usage:")
	fmt.Println("    command [options] [arguments]")

	fmt.Println()
	color.Yellow("Options:")
	for _, command := range c.commands {
		fmt.Print(color.GreenString("    -%s, --%s\t", command.Short, command.Command))
		fmt.Print(command.Help)
	}

	fmt.Println("")
}
