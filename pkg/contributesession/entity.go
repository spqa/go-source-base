package contributesession

import "time"

type Entity struct {
	Id               int64
	OpenTime         time.Time
	ClosureTime      time.Time
	FinalClosureTime time.Time
	ExportedAssets   string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (e *Entity) TableName() string {
	return "contribute_sessions"
}
