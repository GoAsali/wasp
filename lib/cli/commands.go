package cli

import (
	"fmt"
	"github.com/fatih/color"
	"main/lib"
	"strings"
)

type CommandHelp struct {
	Short   string
	Command string
	Help    string
}

type Command struct {
	Command string
	Flags   map[string]interface{}
	Args    []string
}

type CommandCli struct {
	commandsHelp []CommandHelp
	config       lib.Config
	Command      Command
	color        Color
}

func NewCommandCli(args []string) CommandCli {
	commands := []CommandHelp{
		{Help: "Display help for the given command. When no command is given display help for the list command", Short: "h", Command: "help"},
	}

	command := Command{Flags: make(map[string]interface{}), Args: make([]string, 0)}
	if len(args) > 0 {
		command.Command = args[0]
	}

	args = args[1:]
	i := 0
	flag := false
	for _, arg := range args {
		i++

		if !strings.HasPrefix(arg, "-") {
			if !flag {
				command.Args = append(command.Args, arg)
			}
			continue
		}
		flag = true
		if len(args) > i && !strings.HasPrefix(args[i], "-") {
			command.Flags[arg] = args[i]
		} else {
			command.Flags[arg] = true
		}
	}

	return CommandCli{commandsHelp: commands, config: lib.GetConfig(), Command: command}
}

func (c *CommandCli) PrintHelp() {

	fmt.Print("Go Asali Framework ")
	color.Green("%s\n", c.config.Version)
	fmt.Println()
	color.Yellow("Usage:")
	fmt.Println("    command [options] [arguments]")

	fmt.Println()
	color.Yellow("Options:")
	for _, command := range c.commandsHelp {
		fmt.Print(color.GreenString("    -%s, --%s\t", command.Short, command.Command))
		fmt.Print(command.Help)
	}

	fmt.Println("")
}

func (c *CommandCli) Run() {
	executeResult := []bool{NewMakeCommand(c).CheckCommand()}
	for _, result := range executeResult {
		if !result {
			c.color.PrintErrorf("Command %s is not defined.\n", c.Command.Command)
		}
	}
}
