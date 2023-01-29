package template

import "fmt"

type ControllerTemplate struct {
	module string
}

func NewControllerTemplateModule(module string) *ControllerTemplate {
	return &ControllerTemplate{module: module}
}

func NewControllerTemplate() *ControllerTemplate {
	return &ControllerTemplate{module: "main"}
}

func (temp *ControllerTemplate) Crud(controller string) string {
	return fmt.Sprintf(`package %s
	type %s struct{}

	func (c *%s) index() {
		// ToDo: Index controller
	}

	func (c *%s) create() {
		// ToDo: Create controller
	}

	func (c *%s) update() {
		// ToDo: Update controller
	}

	func (c *%s) delete() {
		// ToDo: Delete controller
	}`, temp.module, controller, controller, controller, controller, controller)
}
