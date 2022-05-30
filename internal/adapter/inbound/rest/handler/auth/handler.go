package auth

import (
	"go-boiler-clean/dto"
	"go-boiler-clean/internal/usecase"
	"go-boiler-clean/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type handler struct {
	usecaseAuth usecase.Auth
}

func NewHandler(
	usecaseAuth usecase.Auth,
) *handler {
	return &handler{
		usecaseAuth,
	}
}

func (h *handler) Register(c echo.Context) (err error) {
	ctx := c.Request().Context()
	payload := new(dto.AuthRegister)
	if err = c.Bind(payload); err != nil {
		return response.ErrorBuilder(response.Constant.Error.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		return response.ErrorBuilder(response.Constant.Error.Validation, err).Send(c)
	}

	data, err := h.usecaseAuth.Register(ctx, payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(data).Send(c)
}
