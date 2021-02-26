package systemdata

type Entity struct {
	Key   string
	Value string
}

func (e *Entity) TableName() string {
	return "system_data"
}
