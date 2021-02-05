package router

import (
	"github.com/labstack/echo/v4"
	"mcm-api/pkg/document"
)

type DocumentRouter struct {
	service *document.Service
}

func NewDocumentRouter(service *document.Service) *DocumentRouter {
	return &DocumentRouter{
		service: service,
	}
}

func (router *DocumentRouter) Register(group *echo.Group) {
	group.GET("", router.index)
}

func (router *DocumentRouter) index(echo.Context) error {
	return nil
}
