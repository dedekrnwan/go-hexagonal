package database

import "go-boiler-clean/entity"

var Entity []interface{} = []interface{}{
	&entity.User{},
	&entity.Transaction{},
}
