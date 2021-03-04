package contribution

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
	group.PUT(":id", h.update)
	group.DELETE(":id", h.delete)
}

// @Tags Contributions
// @Summary List contributions
// @Description List contributions
// @Accept  json
// @Produce  json
// @Param params query contribution.IndexQuery false "index query"
// @Success 200 {object} common.PaginateResponse{data=contribution.ContributionRes}
// @Security ApiKeyAuth
// @Router /contributions [get]
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

// @Tags Contributions
// @Summary Show a contribution
// @Description get contribution by ID
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} contribution.ContributionRes
// @Security ApiKeyAuth
// @Router /contributions/{id} [get]
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

// @Tags Contributions
// @Summary Create a contribution
// @Description Create a contribution
// @Accept  json
// @Produce  json
// @Param body body contribution.ContributionCreateReq true "create"
// @Success 200 {object} contribution.ContributionRes
// @Security ApiKeyAuth
// @Router /contributions [post]
func (h *Handler) create(context echo.Context) error {
	body := new(ContributionCreateReq)
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

// @Tags Contributions
// @Summary Update a contribution
// @Description Update a contribution
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Param body body contribution.ContributionUpdateReq true "create"
// @Success 200 {object} contribution.ContributionRes
// @Security ApiKeyAuth
// @Router /contributions/{id} [put]
func (h *Handler) update(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return apperror.HandleError(err, context)
	}
	body := new(ContributionUpdateReq)
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

// @Tags Contributions
// @Summary Delete a contribution
// @Description Delete a contribution
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200
// @Security ApiKeyAuth
// @Router /contributions/{id} [delete]
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
