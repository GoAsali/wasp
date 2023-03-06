package cli

import (
	"fmt"
	MakeService "main/lib/services/make"
	"strings"
)

type CommandMake struct {
	cli *CommandCli
}

func NewMakeCommand(cli *CommandCli) *CommandMake {
	return &CommandMake{
		cli: cli,
	}
}

func (cm *CommandMake) makeController() {
	if len(cm.cli.Command.Args) == 0 {
		cm.cli.logger.PrintError("You must specify a controller name")
		return
	}
	name := cm.cli.Command.Args[0]

	var moduleName string
	flag, hasModuleName := cm.cli.Command.Flags["-m"]
	if hasModuleName {
		moduleName = fmt.Sprintf("%s", flag)
	}

	controller := MakeService.NewMakeController()
	newController, err := controller.CreateNewController(name, moduleName)
	if err != nil {
		cm.cli.logger.PrintError(err.Error())
		return
	}
	cm.cli.logger.PrintSuccess(newController)
}

func (cm *CommandMake) makeMiddleware() {
	if len(cm.cli.Command.Args) == 0 {
		cm.cli.logger.PrintError("You must specify a controller name")
		return
	}
	name := cm.cli.Command.Args[0]

	service := MiddleService
}

func (cm *CommandMake) CheckCommand() bool {

	commands := []string{"controller"}

	parts := strings.Split(cm.cli.Command.Command, ":")
	if parts[0] != "make" {
		return false
	}

	switch parts[1] {
	case "controller":
		cm.makeController()
		break
	default:
		cm.cli.logger.PrintErrorf("Command \"%s\" is not defined.\n", cm.cli.Command.Command)
		cm.cli.logger.PrintError("Did you mean one of these ?\n")
		for _, command := range commands {
			cm.cli.logger.PrintErrorf("\tmake:%s\n", command)
		}
		return true
	}

	return true
}
