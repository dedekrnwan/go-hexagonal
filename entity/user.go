package entity

type User struct {
	BaseEntity

	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email" bson:"email"`
	Phone     string `json:"phone" bson:"phone"`
	IsActive  *bool  `json:"is_active" bson:"is_active"`
	Password  string `json:"password" bson:"password"`

	Transactions []*Transaction `json:"transactions,omitempty" bson:"transactions,omitempty" gorm:"foreignKey:UserId;references:ID"`
}
