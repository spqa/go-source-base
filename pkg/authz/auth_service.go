package authz

type Action string

const (
	Read   Action = "read"
	Create Action = "create"
	Update Action = "update"
	Delete Action = "delete"
)

type ObjectType string

const (
	Faculty           ObjectType = "faculty"
	ContributeSession ObjectType = "contribute_session"
	Contribution      ObjectType = "contribution"
)

type Service struct {
}
