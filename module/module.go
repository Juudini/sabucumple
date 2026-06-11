package module

import "github.com/labstack/echo/v5"

type Module interface {
  Endpoint() string
  Register(g *echo.Group)
}
