package authz

import "github.com/labstack/echo/v4"

type Handler struct {
	service *Service
}

func NewAuthHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(group *echo.Group) {
	group.GET("login", h.login)
}

func (h Handler) login(ctx echo.Context) error {
	return nil
}
