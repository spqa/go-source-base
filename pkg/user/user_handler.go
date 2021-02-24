package user

import (
	"github.com/labstack/echo/v4"
	middleware2 "github.com/labstack/echo/v4/middleware"
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
	group.Use(middleware2.JWT([]byte(h.config.JwtSecret)))
	group.Use(middleware.RequireAuthentication())
	group.GET("", h.index)
	group.GET("/:id", h.getById)
	group.POST("", h.createUser)
}

// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} user.UserResponse
// @Router /accounts/{id} [get]
func (h *Handler) index(echo.Context) error {
	return nil
}

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

func (h Handler) createUser(ctx echo.Context) error {
	loggedInUser := common.GetLoggedInUser(ctx)
	req := &UserCreateReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	userResponse, err := h.service.CreateUser(ctx.Request().Context(), loggedInUser, req)
	if err != nil {
		return apperror.HandleError(err, ctx)
	}
	return ctx.JSON(http.StatusOK, userResponse)
}
