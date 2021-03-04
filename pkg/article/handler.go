package article

import (
	"github.com/labstack/echo/v4"
	"mcm-api/config"
	"mcm-api/pkg/apperror"
	"mcm-api/pkg/middleware"
	"net/http"
	"strconv"
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
	group.GET("/:id", h.getById)
}

// @Tags Articles
// @Summary Show a article
// @Description get article by ID
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} article.ArticleRes
// @Security ApiKeyAuth
// @Router /articles/{id} [get]
func (h *Handler) getById(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}
	result, err := h.service.FindById(context.Request().Context(), id)
	if err != nil {
		return apperror.HandleError(err, context)
	}
	return context.JSON(http.StatusOK, result)
}
