package requests

import (
	"mcm-api/pkg/common"
	"mcm-api/pkg/user"
)

type Query struct {
	Role user.Role `query:"role"`
	common.PaginateQuery
}
