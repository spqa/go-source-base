package internal

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

func (router *UserRouter) index(echo.Context) error {
	return nil
}
