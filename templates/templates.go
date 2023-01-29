package templates

type Template struct {
	Module string
}

func NewTemplate() *Template {
	return &Template{Module: "main"}
}

func NewTemplateInModule(module string) *Template {
	return &Template{Module: module}
}
