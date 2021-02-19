package router

import (
	"github.com/labstack/echo/v4"
	"mcm-api/pkg/response"
	"mcm-api/pkg/user"
	"net/http"
	"strconv"
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
	group.GET("/:id", router.getById)
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

func (router *UserRouter) getById(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return response.NewApiBadRequestError("Id should be string", nil)
	}
	res, err := router.service.FindById(id)
	if err != nil {
		return err
	}
	return context.JSON(http.StatusOK, res)
}
