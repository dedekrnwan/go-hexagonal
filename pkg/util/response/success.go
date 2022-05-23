package response

import (
	"go-boiler-clean/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

type successHelper struct {
	OK Success
}

var successConstant successHelper = successHelper{
	OK: Success{
		Response: successResponse{
			Meta: Meta{
				Success: true,
				Message: "Request successfully proceed",
			},
			Data: nil,
		},
		Code: http.StatusOK,
	},
}

type successResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Success struct {
	Response successResponse `json:"response"`
	Code     int             `json:"code"`
}

func SuccessBuilder(res Success, data interface{}) *Success {
	res.Response.Data = data
	return &res
}

func CustomSuccessBuilder(code int, data interface{}, message string, info *dto.PaginationInfo) *Success {
	return &Success{
		Response: successResponse{
			Meta: Meta{
				Success: true,
				Message: message,
				Info:    info,
			},
			Data: data,
		},
		Code: code,
	}
}

func SuccessResponse(data interface{}) *Success {
	return SuccessBuilder(Constant.Success.OK, data)
}

func (s *Success) Send(c echo.Context) error {
	return c.JSON(s.Code, s.Response)
}
