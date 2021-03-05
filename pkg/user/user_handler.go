package user

import (
	"github.com/labstack/echo/v4"
	"mcm-api/config"
	"mcm-api/pkg/apperror"
	"mcm-api/pkg/common"
	"mcm-api/pkg/middleware"
	"net/http"
	"strconv"
)

type Handler struct {
	config  *config.Config
	service *Service
}

func NewUserHandler(config *config.Config, service *Service) *Handler {
	return &Handler{
		config:  config,
		service: service,
	}
}

func (h *Handler) Register(group *echo.Group) {
	group.Use(middleware.RequireAuthentication(h.config.JwtSecret))
	group.GET("", h.index)
	group.GET("/:id", h.getById)
	group.POST("", h.createUser)
	group.PUT(":id", h.updateUser)
	group.DELETE(":id", h.deleteUser)
}

// @Tags Users
// @Summary List users
// @Description List users
// @Accept  json
// @Produce  json
// @Param params query user.UserIndexQuery false "user index query"
// @Success 200 {object} common.PaginateResponse{data=user.UserResponse}
// @Security ApiKeyAuth
// @Router /users [get]
func (h *Handler) index(ctx echo.Context) error {
	query := new(UserIndexQuery)
	err := ctx.Bind(query)
	if err != nil {
		return err
	}
	loggedInUser, err := common.GetLoggedInUser(ctx.Request().Context())
	if err != nil {
		return err
	}
	paginateRes, err := h.service.Find(ctx.Request().Context(), loggedInUser, query)
	if err != nil {
		return apperror.HandleError(err, ctx)
	}
	return ctx.JSON(http.StatusOK, paginateRes)
}

// @Tags Users
// @Summary Show a user
// @Description get user by ID
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} user.UserResponse
// @Security ApiKeyAuth
// @Router /users/{id} [get]
func (h *Handler) getById(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return apperror.
			New(apperror.ErrInvalid, "Id should be string", err).
			ToResponse(ctx)
	}
	res, err := h.service.FindById(ctx.Request().Context(), id)
	if err != nil {
		return apperror.HandleError(err, ctx)
	}
	return ctx.JSON(http.StatusOK, res)
}

// @Tags Users
// @Summary Create a user
// @Description Create a user
// @Accept  json
// @Produce  json
// @Param user body user.UserCreateReq true "create user"
// @Success 200 {object} user.UserResponse
// @Security ApiKeyAuth
// @Router /users [post]
func (h Handler) createUser(ctx echo.Context) error {
	loggedInUser, err := common.GetLoggedInUser(ctx.Request().Context())
	if err != nil {
		return apperror.HandleError(err, ctx)
	}
	req := &UserCreateReq{}
	err = ctx.Bind(req)
	if err != nil {
		return err
	}
	userResponse, err := h.service.CreateUser(ctx.Request().Context(), loggedInUser, req)
	if err != nil {
		return apperror.HandleError(err, ctx)
	}
	return ctx.JSON(http.StatusOK, userResponse)
}

// @Tags Users
// @Summary Update a user
// @Description Update a user
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body user.UserCreateReq true "create user"
// @Success 200 {object} user.UserResponse
// @Security ApiKeyAuth
// @Router /users/{id} [put]
func (h Handler) updateUser(ctx echo.Context) error {
	return nil
}

// @Tags Users
// @Summary Delete a user
// @Description Delete a user
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200
// @Security ApiKeyAuth
// @Router /users/{id} [delete]
func (h Handler) deleteUser(ctx echo.Context) error {
	return nil
}
