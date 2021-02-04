package internal

import "github.com/labstack/echo/v4"

func RegisterUserRouter(group *echo.Group) {
	group.GET("", index)
}

func index(echo.Context) error {
	return nil
}
