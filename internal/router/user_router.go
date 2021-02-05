package router

import (
	"github.com/labstack/echo/v4"
	"mcm-api/pkg/user"
)

type UserRouter struct {
	service *user.Service
}

func NewUserRouter(service *user.Service) *UserRouter {
	return &UserRouter{
		service: service,
	}
}

func (router *UserRouter) Register(group *echo.Group) {
	group.GET("", router.index)
}

// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} responses.UserResponse
// @Router /accounts/{id} [get]
func (router *UserRouter) index(echo.Context) error {
	return nil
}
