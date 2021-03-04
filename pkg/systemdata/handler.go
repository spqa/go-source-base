package systemdata

import (
	"github.com/labstack/echo/v4"
	"mcm-api/config"
	"mcm-api/pkg/apperror"
	"mcm-api/pkg/middleware"
	"net/http"
)

type Handler struct {
	config  *config.Config
	service *Service
}

func NewHandler(config *config.Config, service *Service) *Handler {
	return &Handler{
		config:  config,
		service: service,
	}
}

func (h *Handler) Register(group *echo.Group) {
	group.Use(middleware.RequireAuthentication(h.config.JwtSecret))
	group.GET("", h.index)
	group.PUT(":id", h.update)
}

// @Tags System Data
// @Summary Get System Data
// @Description Get System Data
// @Accept  json
// @Produce  json
// @Success 200 {object} systemdata.DataRes
// @Security ApiKeyAuth
// @Router /system-data [get]
func (h *Handler) index(context echo.Context) error {
	paginateResponse, err := h.service.Find(context.Request().Context())
	if err != nil {
		return apperror.HandleError(err, context)
	}
	return context.JSON(http.StatusOK, paginateResponse)
}

// @Tags System Data
// @Summary Update system data
// @Description Update system data
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param body body systemdata.DataUpdateReq true "create"
// @Success 200
// @Security ApiKeyAuth
// @Router /system-data/{id} [put]
func (h *Handler) update(context echo.Context) error {
	id := context.Param("id")
	body := new(DataUpdateReq)
	err := context.Bind(body)
	if err != nil {
		return apperror.HandleError(err, context)
	}
	err = h.service.Update(context.Request().Context(), id, body)
	if err != nil {
		return apperror.HandleError(err, context)
	}
	return context.NoContent(http.StatusOK)
}
