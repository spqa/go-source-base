package user

type Entity struct {
	Id   int
	Name string
}

func (e *Entity) TableName() string {
	return "users"
}
