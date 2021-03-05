package user

import (
	"mcm-api/pkg/common"
	"time"
)

type Entity struct {
	Id        int
	Name      string
	Email     string
	Password  string
	FacultyId *int
	Role      common.Role
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (e *Entity) TableName() string {
	return "users"
}
