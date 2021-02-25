package faculty

import "time"

type Entity struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (e *Entity) TableName() string {
	return "faculties"
}
