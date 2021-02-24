package common

import "github.com/labstack/echo/v4"

type Role string

const (
	Administrator        Role = "admin"
	MarketingManager     Role = "marketing_manager"
	MarketingCoordinator Role = "marketing_coordinator"
	Student              Role = "student"
	Guest                Role = "guest"
)

type LoggedInUser struct {
	Id    int
	Email string
	Name  string
	Role  Role
}

func GetLoggedInUser(ctx echo.Context) *LoggedInUser {
	return ctx.Get("user").(*LoggedInUser)
}
