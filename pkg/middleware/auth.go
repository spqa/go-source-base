package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"mcm-api/pkg/apperror"
	"mcm-api/pkg/common"
	"mcm-api/pkg/log"
	"strconv"
)

func RequireAuthentication(jwtSecret string) echo.MiddlewareFunc {
	return Compose(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(jwtSecret),
		ErrorHandlerWithContext: func(err error, context echo.Context) error {
			log.Logger.Debug("JWT error", zap.Error(err))
			appError := apperror.New(apperror.ErrUnauthorized, "invalid token", nil)
			return apperror.HandleError(appError, context)
		},
	}), requireAuthentication())
}

func requireAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			u := c.Get("user").(*jwt.Token)
			claims := u.Claims.(jwt.MapClaims)
			id, _ := strconv.Atoi(claims["sub"].(string))
			var facultyId int
			if claims["facultyId"] != nil {
				facultyId, _ = strconv.Atoi(claims["facultyId"].(string))
			}
			common.SetLoggedInUser(c, &common.LoggedInUser{
				Id:        id,
				Email:     claims["email"].(string),
				Name:      claims["name"].(string),
				Role:      common.Role(claims["role"].(string)),
				FacultyId: &facultyId,
			})
			return next(c)
		}
	}
}

func Compose(middlewares ...echo.MiddlewareFunc) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			h := next
			for i := len(middlewares) - 1; i >= 0; i-- {
				h = middlewares[i](h)
			}
			return h(c)
		}
	}
}
