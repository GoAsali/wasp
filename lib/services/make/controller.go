package make

import (
	template "main/templates/controllers"
)

type Controller struct {
	name string
	make *Make
}

func NewMakeController() *Controller {
	make := CreateNewMakeService("controllers", "controller")

	return &Controller{
		make: make,
	}
}

func (c *Controller) CreateNewController(name string, moduleName string) (string, error) {
	var temp *template.ControllerTemplate
	if moduleName == "" {
		temp = template.NewControllerTemplate()
	} else {
		temp = template.NewControllerTemplateModule(moduleName)
	}

	return c.make.createFile(name, temp.Crud(name))
}
