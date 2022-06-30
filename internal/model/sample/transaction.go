package sample

import (
	"go-boiler-clean/internal/model/common"
)

type (
	Transaction interface {
		common.Common[TransactionEntity, TransactionEntity]
	}

	TransactionEntity struct {
		CommonEntity

		Type        string  `json:"type" bson:"type"`
		Amount      float64 `json:"amount" bson:"amount"`
		ReferenceNo string  `json:"reference_no" bson:"reference_no"`
		Notes       string  `json:"notes" bson:"notes"`

		UserId *int `json:"user_id" bson:"user_id"`

		User *User `json:"user,omitempty" bson:"user,omitempty"`
	}
)
