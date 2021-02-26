package article

import (
	"gorm.io/datatypes"
	"time"
)

type Status string

const (
	Accepted  Status = "accepted"
	Rejected  Status = "rejected"
	Reviewing Status = "reviewing"
)

type Entity struct {
	Id                  int64
	UserId              int64
	ContributeSessionId int64
	ArticleId           int64
	Images              datatypes.JSON
	Status              Status
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func (e *Entity) TableName() string {
	return "contributions"
}
