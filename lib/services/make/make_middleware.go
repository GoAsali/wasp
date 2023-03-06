package make

import "main/templates/middlewares"

type MakeMiddleware struct {
	make Make
}

func CreateNewMiddlewareService() Make {
	return Make{ModuleType: "middleware", folder: "middlewares"}
}

func (m *MakeMiddleware) CreateNewMiddleware(name string) (string, error) {
	temp = middlewares.CreateNewMiddlewareTemplate()
	temp.
		m.make.createFile(name)
}
