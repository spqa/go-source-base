package apperror

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"mcm-api/pkg/log"
	"net/http"
)

type AppErrCode string

const (
	ErrConflict     AppErrCode = "conflict"
	ErrInternal     AppErrCode = "internal"
	ErrInvalid      AppErrCode = "invalid"
	ErrNotFound     AppErrCode = "not_found"
	ErrForbidden    AppErrCode = "forbidden"
	ErrUnauthorized AppErrCode = "unauthorized"
)

type appError struct {
	Code    AppErrCode
	Message string
	Err     error
	Data    interface{}
}

type appErrorRes struct {
	Message string      `json:"message"`
	Code    AppErrCode  `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}

func (a *appError) Error() string {
	return a.Message
}

func (a *appError) WithData(data interface{}) *appError {
	a.Data = data
	return a
}

func (a *appError) ToResponse(ctx echo.Context) error {
	switch a.Code {
	case ErrNotFound:
		return ctx.JSON(http.StatusNotFound, appErrorRes{
			Message: valueOrDefault(a.Message, "not found"),
			Code:    a.Code,
			Data:    a.Data,
		})
	case ErrConflict:
		return ctx.JSON(http.StatusConflict, appErrorRes{
			Message: valueOrDefault(a.Message, "conflict"),
			Code:    a.Code,
			Data:    a.Data,
		})
	case ErrForbidden:
		return ctx.JSON(http.StatusForbidden, appErrorRes{
			Message: valueOrDefault(a.Message, "forbidden"),
			Code:    a.Code,
			Data:    a.Data,
		})
	case ErrInvalid:
		return ctx.JSON(http.StatusBadRequest, appErrorRes{
			Message: valueOrDefault(a.Message, "invalid"),
			Code:    a.Code,
			Data:    a.Data,
		})
	case ErrInternal:
		return ctx.JSON(http.StatusInternalServerError, appErrorRes{
			Message: valueOrDefault(a.Message, "internal server error"),
			Code:    a.Code,
			Data:    a.Data,
		})
	case ErrUnauthorized:
		return ctx.JSON(http.StatusUnauthorized, appErrorRes{
			Message: valueOrDefault(a.Message, "unauthorized"),
			Code:    a.Code,
			Data:    a.Data,
		})
	default:
		return ctx.JSON(http.StatusInternalServerError, appErrorRes{
			Message: valueOrDefault(a.Message, "internal server error"),
			Code:    a.Code,
			Data:    a.Data,
		})
	}
}

func valueOrDefault(str string, def string) string {
	if str == "" {
		return def
	}
	return str
}

func Is(err error, code AppErrCode) bool {
	if v, ok := err.(*appError); ok {
		return v.Code == code
	}
	return false
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
	case validation.Errors:
		return New(ErrInvalid, "", err).WithData(err).ToResponse(ctx)
	case *echo.HTTPError:
		return New(ErrInvalid, e.Unwrap().Error(), e.Unwrap()).ToResponse(ctx)
	default:
		log.Logger.Error("unhandled error", zap.Error(err))
		return ctx.JSON(http.StatusInternalServerError, appErrorRes{
			Message: "Internal server error",
			Code:    ErrInternal,
		})
	}
}
