package transaction

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) Get(c echo.Context) error {
	return c.String(http.StatusOK, "testing sample")
}
