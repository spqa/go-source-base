package comment

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
	group.PUT("/:id", h.update)
	group.DELETE("/:id", h.delete)
}

// @Tags Comments
// @Summary List comments
// @Description List comments
// @Accept  json
// @Produce  json
// @Param params query comment.IndexQuery false "index query"
// @Success 200 {object} common.PaginateResponse{data=comment.CommentRes}
// @Security ApiKeyAuth
// @Router /comments [get]
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

// @Tags Comments
// @Summary Show a comment
// @Description get comment by ID
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} comment.CommentRes
// @Security ApiKeyAuth
// @Router /comments/{id} [get]
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

// @Tags Comments
// @Summary Create a comment
// @Description Create a comment
// @Accept  json
// @Produce  json
// @Param body body comment.CommentCreateReq true "create"
// @Success 200 {object} comment.CommentRes
// @Security ApiKeyAuth
// @Router /comments [post]
func (h *Handler) create(context echo.Context) error {
	body := new(CommentCreateReq)
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

// @Tags Comments
// @Summary Update a comment
// @Description Update a comment
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Param body body comment.CommentUpdateReq true "update"
// @Success 200 {object} comment.CommentRes
// @Security ApiKeyAuth
// @Router /comments/{id} [put]
func (h *Handler) update(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return apperror.HandleError(err, context)
	}
	body := new(CommentUpdateReq)
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

// @Tags Comments
// @Summary Delete a comment
// @Description Delete a comment
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200
// @Security ApiKeyAuth
// @Router /comments/{id} [delete]
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
