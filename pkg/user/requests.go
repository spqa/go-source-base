package user

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"mcm-api/pkg/common"
)

type UserIndexQuery struct {
	Role common.Role `query:"role" enums:"admin,marketing_manager,marketing_coordinator,student,guest"`
	common.PaginateQuery
}

type UserCreateReq struct {
	Name      string      `json:"name"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	Role      common.Role `json:"role"`
	FacultyId int         `json:"facultyId"`
}

func (c *UserCreateReq) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Name, validation.Required, validation.Length(5, 50)),
		validation.Field(&c.Email, validation.Required, is.Email),
		validation.Field(&c.Password, validation.Required, validation.Length(5, 50)),
		validation.Field(&c.Role, validation.Required, validation.In(
			common.Guest, common.Student, common.MarketingCoordinator, common.MarketingManager, common.Administrator)),
		validation.Field(&c.FacultyId, validation.Required.When(isRoleRequiredFaculty(c.Role)), is.Int),
	)
}

func isRoleRequiredFaculty(role common.Role) bool {
	switch role {
	case common.Administrator:
		return false
	case common.MarketingManager:
		return false
	case common.MarketingCoordinator:
		return true
	case common.Student:
		return true
	case common.Guest:
		return true
	default:
		return true
	}
}
