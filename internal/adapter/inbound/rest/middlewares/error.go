package middlewares

import (
	"go-boiler-clean/pkg/util/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	var errCustom *response.Error

	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	switch report.Code {
	case http.StatusNotFound:
		errCustom = response.ErrorBuilder(response.Constant.Error.RouteNotFound, err)
	default:
		errCustom = response.ErrorBuilder(response.Constant.Error.InternalServerError, err)
	}

	response.ErrorResponse(errCustom).Send(c)
}
