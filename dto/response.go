package dto

type (
	ResponseSingle[T any] struct {
		Data T `json:"data" bson:"data"`
	}
	ResponseMany[T any] struct {
		Data           []T            `json:"data" bson:"data"`
		PaginationInfo PaginationInfo `json:"pagination_info" bson:"pagination_info"`
	}
)
