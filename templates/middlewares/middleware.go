package middlewares

import (
	"fmt"
	"main/templates"
)

type MiddlewareTemplate struct {
	template templates.Template
}

func CreateNewMiddlewareTemplate() *MiddlewareTemplate {
	module := templates.Template{Module: "middlewares"}
	return &MiddlewareTemplate{
		template: module,
	}
}

func (temp MiddlewareTemplate) CreateMiddleware(name string) string {
	return fmt.Sprintf(`package %s

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func %s() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
`, temp.template.Module, name)
}
