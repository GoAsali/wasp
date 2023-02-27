package cli

import (
	"fmt"
	"log"
	template "main/templates/controllers"
	"os"
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

	// Check controller folder was exists, if not create it
	if _, err := os.Stat("controllers"); os.IsNotExist(err) {
		err := os.Mkdir("controllers", 0755)
		if err != nil {
			if os.IsPermission(err) {
				cm.cli.logger.PrintError("You need permission to current folder to create controllers\n")
				return
			}
			cm.cli.logger.PrintError("Could not create controllers")
			return
		}
	}

	fileName := fmt.Sprintf("controllers/%s.go", name)

	if _, err := os.Stat(fileName); err == nil {
		cm.cli.logger.PrintErrorf("Controller %s already exists\n", name)
		return
	}

	file, e := os.Create(fileName)
	if e != nil {
		log.Fatal(e)
	}

	//Module name
	var moduleName string
	flag, hasModuleName := cm.cli.Command.Flags["-m"]
	if hasModuleName {
		moduleName = fmt.Sprintf("%s", flag)
	}

	var temp *template.ControllerTemplate
	if moduleName == "" {
		temp = template.NewControllerTemplate()
	} else {
		temp = template.NewControllerTemplateModule(moduleName)
	}

	_, err := file.WriteString(temp.Crud(name))
	if err != nil {
		return
	}
	err = file.Close()
	if err != nil {
		cm.cli.logger.PrintErrorf("Error during writing template", "")
	}

	cm.cli.logger.PrintSuccess("Controller created successfully\n")
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
