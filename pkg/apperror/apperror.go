package apperror

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"mcm-api/pkg/log"
	"net/http"
)

type AppErrCode string

const (
	ErrConflict  AppErrCode = "conflict"
	ErrInternal  AppErrCode = "internal"
	ErrInvalid   AppErrCode = "invalid"
	ErrNotFound  AppErrCode = "not_found"
	ErrForbidden AppErrCode = "forbidden"
)

type appError struct {
	Code    AppErrCode
	Message string
	Err     error
}

type appErrorRes struct {
	Message string     `json:"message"`
	Code    AppErrCode `json:"code"`
}

func (a *appError) Error() string {
	return a.Message
}

func (a *appError) ToResponse(ctx echo.Context) error {
	switch a.Code {
	case ErrNotFound:
		return ctx.JSON(http.StatusNotFound, appErrorRes{
			Message: valueOrDefault(a.Message, "not found"),
			Code:    a.Code,
		})
	case ErrConflict:
		return ctx.JSON(http.StatusConflict, appErrorRes{
			Message: valueOrDefault(a.Message, "conflict"),
			Code:    a.Code,
		})
	case ErrForbidden:
		return ctx.JSON(http.StatusForbidden, appErrorRes{
			Message: valueOrDefault(a.Message, "forbidden"),
			Code:    a.Code,
		})
	case ErrInvalid:
		return ctx.JSON(http.StatusBadRequest, appErrorRes{
			Message: valueOrDefault(a.Message, "invalid"),
			Code:    a.Code,
		})
	case ErrInternal:
		return ctx.JSON(http.StatusInternalServerError, appErrorRes{
			Message: valueOrDefault(a.Message, "internal server error"),
			Code:    a.Code,
		})
	default:
		return ctx.JSON(http.StatusInternalServerError, appErrorRes{
			Message: valueOrDefault(a.Message, "internal server error"),
			Code:    a.Code,
		})
	}
}

func valueOrDefault(str string, def string) string {
	if str == "" {
		return def
	}
	return str
}

func New(code AppErrCode, message string, err error) *appError {
	return &appError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func HandleError(err error, ctx echo.Context) error {
	switch e := err.(type) {
	case *appError:
		return e.ToResponse(ctx)
	default:
		log.Logger.Error("unhandled error", zap.Error(err))
		return ctx.JSON(http.StatusInternalServerError, appErrorRes{
			Message: "Internal server error",
			Code:    ErrInternal,
		})
	}
}
