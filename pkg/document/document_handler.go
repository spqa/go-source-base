package document

import (
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *Service
}

func NewDocumentHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(group *echo.Group) {
	group.GET("", h.index)
}

func (h *Handler) index(echo.Context) error {
	return nil
}
