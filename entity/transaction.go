package entity

type Transaction struct {
	BaseEntity

	Type        string  `json:"type" bson:"type" sql:"type:ENUM('debit','kredit')"`
	Amount      float64 `json:"amount" bson:"amount"`
	ReferenceNo string  `json:"reference_no" bson:"reference_no"`
	Notes       string  `json:"notes" bson:"notes"`

	UserId *int `json:"user_id" bson:"user_id"`

	User *User `json:"user,omitempty" bson:"user,omitempty" gorm:"foreignKey:UserId"`
}
