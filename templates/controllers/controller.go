package template

import (
	"fmt"
	"main/templates"
)

type ControllerTemplate struct {
	template *templates.Template
}

func newControllerTemplate(temp *templates.Template) *ControllerTemplate {
	return &ControllerTemplate{template: temp}

}
func NewControllerTemplateModule(module string) *ControllerTemplate {
	return newControllerTemplate(templates.NewTemplateInModule(module))
}

func NewControllerTemplate() *ControllerTemplate {
	return &ControllerTemplate{template: templates.NewTemplate()}
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
}`, temp.template.Module, controller, controller, controller, controller, controller)
}
