package make

import "main/templates/middlewares"

type Middleware struct {
	make *Make
}

func CreateNewMiddlewareService() Middleware {
	return Middleware{
		make: CreateNewMakeService("middlewares", "middleware"),
	}
}

func (m *Middleware) CreateNewMiddleware(name string) (string, error) {
	temp := middlewares.CreateNewMiddlewareTemplate()
	return m.make.createFile(name, temp.CreateMiddleware(name))
}
