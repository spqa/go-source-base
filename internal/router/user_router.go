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
	group.GET("/id", router.getById)
	group.POST("", router.create)
	group.PUT("/:id", router.update)
	group.DELETE("/:id", router.delete)
}

// @Tags User
// @Summary Get user list
// @Description Get user list
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} responses.UserResponse
// @Router /users [get]
func (router *UserRouter) index(ctx echo.Context) error {
	return nil
}

// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} responses.UserResponse
// @Router /accounts/{id} [get]
func (router *UserRouter) getById(ctx echo.Context) error {
	return nil
}

func (router *UserRouter) create(ctx echo.Context) error {
	return nil
}

func (router *UserRouter) update(ctx echo.Context) error {
	return nil
}

func (router *UserRouter) delete(ctx echo.Context) error {
	return nil
}
