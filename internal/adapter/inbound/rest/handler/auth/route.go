package auth

import "github.com/labstack/echo/v4"

func (h *handler) Route(g *echo.Group) {
	g.POST("/registers", h.Register)
}