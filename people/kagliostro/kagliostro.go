package kagliostro

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type Module struct{}

func New() Module {
	return Module{}
}

func (Module) Endpoint() string {
	return "kagliostro"
}

func (Module) Register(g *echo.Group) {
	g.GET("/", home)
}

func home(c *echo.Context) error {
	return c.HTML(http.StatusOK, `
		<h1>kagli</h1>
		<p>Aquí ejemplo puede hacer lo que quiera.</p>
		`)
}
