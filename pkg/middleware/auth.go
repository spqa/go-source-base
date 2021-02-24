package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"mcm-api/pkg/common"
	"strconv"
)

//func RequireAuthentication(jwtSecret string) echo.MiddlewareFunc {
//	jwtMiddleware := middleware.JWT([]byte(jwtSecret))
//	authMiddleware := requireAuthentication()
//	return func(next echo.HandlerFunc) echo.HandlerFunc {
//		return func(c echo.Context) error {
//			h := next
//			jwtMiddleware(h)
//			authMiddleware(h)
//			return h(c)
//		}
//	}
//}

func RequireAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Println("haha")
			u := c.Get("user").(*jwt.Token)
			claims := u.Claims.(jwt.MapClaims)
			id, _ := strconv.Atoi(claims["sub"].(string))
			c.Set("user", &common.LoggedInUser{
				Id:    id,
				Email: claims["email"].(string),
				Name:  claims["name"].(string),
				Role:  common.Role(claims["role"].(string)),
			})
			return next(c)
		}
	}
}
