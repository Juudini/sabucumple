package module

import "github.com/labstack/echo/v4"

type Module interface {
	Endpoint() string
	Register(g *echo.Group)
}
