package user

type Role int

const (
	Administrator        Role = 1
	MarketingManager     Role = 2
	MarketingCoordinator Role = 3
	Student              Role = 4
	Guest                Role = 5
)

type Entity struct {
	Id   int
	Name string
	Role Role
}

func (e *Entity) TableName() string {
	return "users"
}
