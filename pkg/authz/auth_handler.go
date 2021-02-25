package authz

import (
	"github.com/labstack/echo/v4"
	"mcm-api/pkg/apperror"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewAuthHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(group *echo.Group) {
	group.POST("/login", h.login)
}

// @Tags Auth
// @Summary Login
// @Description Login
// @Accept  json
// @Produce  json
// @Param body body authz.LoginRequest true "login req"
// @Success 200 {object} authz.LoginResponse
// @Router /auth/login [post]
func (h Handler) login(ctx echo.Context) error {
	req := new(LoginRequest)
	err := ctx.Bind(&req)
	if err != nil {
		return err
	}
	loginResponse, err := h.service.Login(ctx.Request().Context(), req)
	if err != nil {
		return apperror.HandleError(err, ctx)
	}
	return ctx.JSON(http.StatusOK, loginResponse)
}
