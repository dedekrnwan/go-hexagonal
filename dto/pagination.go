package dto

type (
	Pagination struct {
		Page  *int `query:"page" json:"page" bson:"page"`
		Limit *int `query:"limit" json:"limit" bson:"limit"`
	}

	PaginationInfo struct {
		Pagination
		Count     int64 `json:"count" bson:"count"`
		TotalPage int64 `json:"total_page" bson:"total_page"`
	}
)
