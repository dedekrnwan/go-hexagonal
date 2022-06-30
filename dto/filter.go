package dto

type (
	Filter struct {
		Field    string      `json:"field" bson:"field"`
		Operator string      `json:"operator" bson:"operator"`
		Value    interface{} `json:"value" bson:"value"`
	}
)
