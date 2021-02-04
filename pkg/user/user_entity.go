package user

import "gorm.io/gorm"

type Entity struct {
	gorm.Model
	Id   int
	Name string
}
