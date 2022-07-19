package dto

import "time"

type (
	User struct {
		ID int `json:"id" gorm:"primaryKey;autoIncrement;" param:"id"`

		CreatedAt time.Time `json:"created_at"`
		CreatedBy *int      `json:"created_by"`

		ModifiedAt time.Time `json:"modified_at"`
		ModifiedBy *int      `json:"modified_by"`

		DeletedAt time.Time `json:"-" gorm:"index"`
		DeletedBy *int      `json:"deleted_by"`

		FirstName string `json:"first_name" bson:"first_name"`
		LastName  string `json:"last_name" bson:"last_name"`
		Email     string `json:"email" bson:"email"`
		Phone     string `json:"phone" bson:"phone"`
		IsActive  *bool  `json:"is_active" bson:"is_active"`
		Password  string `json:"-" bson:"password"`

		Transactions []Transaction `json:"transactions,omitempty" bson:"transactions,omitempty"`
	}
)
