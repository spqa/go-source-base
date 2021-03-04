package contributesession

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
	group.GET("", h.index)
	group.GET("/:id", h.getById)
	group.POST("", h.create)
	group.POST("/:id/export", h.create)
	group.PUT("/:id", h.update)
	group.DELETE("/:id", h.delete)
}

// @Tags Contribute Sessions
// @Summary List Contribute Sessions
// @Description List Contribute Sessions
// @Accept  json
// @Produce  json
// @Param params query contributesession.IndexQuery false "index query"
// @Success 200 {object} common.PaginateResponse{data=contributesession.SessionRes}
// @Security ApiKeyAuth
// @Router /contribute-sessions [get]
func (h *Handler) index(context echo.Context) error {
	query := new(IndexQuery)
	err := context.Bind(query)
	if err != nil {
		return apperror.HandleError(err, context)
	}
	paginateResponse, err := h.service.Find(context.Request().Context(), query)
	if err != nil {
		return apperror.HandleError(err, context)
	}
	return context.JSON(http.StatusOK, paginateResponse)
}

// @Tags Contribute Sessions
// @Summary Show a Contribute Session
// @Description get Contribute Session by ID
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} contributesession.SessionRes
// @Security ApiKeyAuth
// @Router /contribute-sessions/{id} [get]
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

// @Tags Contribute Sessions
// @Summary Create a Contribute Session
// @Description Create a Contribute Session
// @Accept  json
// @Produce  json
// @Param body body contributesession.SessionCreateReq true "create"
// @Success 200 {object} contributesession.SessionRes
// @Security ApiKeyAuth
// @Router /contribute-sessions [post]
func (h *Handler) create(context echo.Context) error {
	body := new(SessionCreateReq)
	err := context.Bind(body)
	if err != nil {
		return apperror.HandleError(err, context)
	}
	result, err := h.service.Create(context.Request().Context(), body)
	if err != nil {
		return apperror.HandleError(err, context)
	}
	return context.JSON(http.StatusOK, result)
}

// @Tags Contribute Sessions
// @Summary Update a Contribute Session
// @Description Update a Contribute Session
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Param body body contributesession.SessionUpdateReq true "create"
// @Success 200 {object} contributesession.SessionRes
// @Security ApiKeyAuth
// @Router /contribute-sessions/{id} [put]
func (h *Handler) update(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return apperror.HandleError(err, context)
	}
	body := new(SessionUpdateReq)
	err = context.Bind(body)
	if err != nil {
		return apperror.HandleError(err, context)
	}
	result, err := h.service.Update(context.Request().Context(), id, body)
	if err != nil {
		return apperror.HandleError(err, context)
	}
	return context.JSON(http.StatusOK, result)
}

// @Tags Contribute Sessions
// @Summary Delete a Contribute Session
// @Description Delete a Contribute Session
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200
// @Security ApiKeyAuth
// @Router /contribute-sessions/{id} [delete]
func (h *Handler) delete(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return apperror.HandleError(err, context)
	}
	err = h.service.Delete(context.Request().Context(), id)
	if err != nil {
		return apperror.HandleError(err, context)
	}
	return context.NoContent(http.StatusNoContent)
}
