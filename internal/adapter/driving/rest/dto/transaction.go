package dto

import "time"

type (
	Transaction struct {
		ID int `json:"id" param:"id"`

		CreatedAt time.Time `json:"created_at"`
		CreatedBy *int      `json:"created_by"`

		ModifiedAt time.Time `json:"modified_at"`
		ModifiedBy *int      `json:"modified_by"`

		DeletedAt time.Time `json:"-" gorm:"index"`
		DeletedBy *int      `json:"deleted_by"`

		Type        string  `json:"type" bson:"type"`
		Amount      float64 `json:"amount" bson:"amount"`
		ReferenceNo string  `json:"reference_no" bson:"reference_no"`
		Notes       string  `json:"notes" bson:"notes"`

		UserId *int `json:"user_id" bson:"user_id"`

		User *User `json:"user,omitempty" bson:"user,omitempty"`
	}
)
