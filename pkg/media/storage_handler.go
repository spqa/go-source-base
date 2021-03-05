package media

import (
	"errors"
	"github.com/labstack/echo/v4"
	"mcm-api/config"
	"mcm-api/pkg/apperror"
	"mcm-api/pkg/common"
	"mcm-api/pkg/middleware"
	"net/http"
)

type Handler struct {
	config  *config.Config
	service Service
}

func NewHandler(config *config.Config, service Service) *Handler {
	return &Handler{
		config:  config,
		service: service,
	}
}

func (h *Handler) Register(group *echo.Group) {
	group.Use(middleware.RequireAuthentication(h.config.JwtSecret))
	group.POST("/upload", h.upload)
}

// @Tags Storage
// @Summary Upload file
// @Description Upload file
// @Accept  multipart/form-data
// @Produce  json
// @Param params query media.UploadQuery true "query"
// @Param file formData file true "upload file"
// @Success 200 {object} media.UploadResult
// @Security ApiKeyAuth
// @Router /storage/upload [post]
func (h *Handler) upload(ctx echo.Context) error {
	query := new(UploadQuery)
	_ = ctx.Bind(query)
	user, err := common.GetLoggedInUser(ctx.Request().Context())
	if err != nil {
		return apperror.HandleError(err, ctx)
	}
	file, err := ctx.FormFile("file")
	if err != nil {
		if errors.Is(err, http.ErrMissingFile) {
			return apperror.HandleError(apperror.New(apperror.ErrInvalid, "missing file", err), ctx)
		}
		return apperror.HandleError(err, ctx)
	}
	open, err := file.Open()
	if err != nil {
		return apperror.HandleError(err, ctx)
	}
	defer func() {
		_ = open.Close()
	}()
	var result *UploadResult
	switch query.Type {
	case Document:
		result, err = h.service.UploadDocumentOriginal(ctx.Request().Context(), &FileUploadOriginalReq{
			File: open,
			Size: file.Size,
			Name: file.Filename,
			User: user,
		})
		break
	case Image:
		result, err = h.service.UploadImage(ctx.Request().Context(), &FileUploadOriginalReq{
			File: open,
			Size: file.Size,
			Name: file.Filename,
			User: user,
		})
	default:
		return apperror.HandleError(apperror.New(apperror.ErrInvalid, "unknown upload type", nil), ctx)
	}
	if err != nil {
		return apperror.HandleError(err, ctx)
	}
	return ctx.JSON(http.StatusOK, result)
}
