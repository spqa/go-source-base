package user

import (
	"github.com/labstack/echo/v4"
	"mcm-api/pkg/apperror"
	"net/http"
	"strconv"
)

type Handler struct {
	service *Service
}

func NewUserHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(group *echo.Group) {
	group.GET("", h.index)
	group.GET("/:id", h.getById)
}

// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} responses.UserResponse
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
