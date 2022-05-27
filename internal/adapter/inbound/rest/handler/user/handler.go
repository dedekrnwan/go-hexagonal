package user

import (
	"go-boiler-clean/entity"
	dtoRest "go-boiler-clean/internal/adapter/inbound/rest/dto"
	"go-boiler-clean/internal/usecase"
	"go-boiler-clean/pkg/util/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	usecaseUser usecase.User
}

func NewHandler(
	usecaseUser usecase.User,
) *handler {
	return &handler{
		usecaseUser,
	}
}

func (h *handler) Get(c echo.Context) error {
	ctx := c.Request().Context()
	payload := dtoRest.NewHttpQuery(c.Request(), entity.User{})
	if err := c.Bind(payload); err != nil {
		return c.String(http.StatusBadRequest, "testing user failed")
	}
	payload.BindFilters()
	// if err := c.Validate(payload); err != nil {
	// 	fmt.Println(err.Error())
	// 	return c.String(http.StatusBadRequest, "testing user failed")
	// }
	// for _, v := range payload.Filters {
	// 	fmt.Printf("%s %s %s\n", v.Field, v.Operator, v.Value)
	// }
	// for _, v := range payload.Ascending {
	// 	fmt.Printf("%s n", v)
	// }
	// fmt.Println(payload.Search)
	// fmt.Println(payload.Ascending)
	data, info, err := h.usecaseUser.Find(ctx, payload.Search, payload.Filters, payload.Ascending, payload.Descending, payload.Pagination)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.CustomSuccessBuilder(response.Constant.Success.OK.Code, data, "Data has been retrieve", info).Send(c)
}
