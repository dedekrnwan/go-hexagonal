package dto

type (
	Filter struct {
		Field    string `json:"field" bson:"field"`
		Operator string `json:"operator" bson:"operator"`
		Value    string `json:"value" bson:"value"`
	}
)
