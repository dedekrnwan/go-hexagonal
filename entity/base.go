package entity

import (
	"time"
)

type Entities interface {
	User | Transaction
}

type BaseEntity struct {
	ID int `json:"id" gorm:"primaryKey;autoIncrement;" param:"id" swaggerignore:"true"`

	CreatedAt time.Time `json:"created_at" swaggerignore:"true"`
	CreatedBy *int      `json:"created_by" swaggerignore:"true"`

	ModifiedAt time.Time `json:"modified_at" swaggerignore:"true"`
	ModifiedBy *int      `json:"modified_by" swaggerignore:"true"`

	DeletedAt time.Time `json:"-" gorm:"index" swaggerignore:"true"`
	DeletedBy *int      `json:"deleted_by" swaggerignore:"true"`
}
