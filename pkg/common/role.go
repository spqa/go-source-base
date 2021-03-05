package common

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"mcm-api/pkg/apperror"
)

type Role string

const (
	Administrator        Role = "admin"
	MarketingManager     Role = "marketing_manager"
	MarketingCoordinator Role = "marketing_coordinator"
	Student              Role = "student"
	Guest                Role = "guest"
)

type LoggedInUser struct {
	Id        int
	Email     string
	Name      string
	Role      Role
	FacultyId *int
}

const contextKey = "user"

func GetLoggedInUser(ctx context.Context) (*LoggedInUser, error) {
	fmt.Println(ctx.Value(contextKey))
	if v, oke := ctx.Value(contextKey).(*LoggedInUser); oke {
		if oke {
			return v, nil
		}
	}
	return nil, apperror.New(apperror.ErrUnauthorized, "cant get logged in user", nil)
}

func SetLoggedInUser(ctx echo.Context, user *LoggedInUser) {
	ctx.SetRequest(ctx.Request().Clone(context.WithValue(ctx.Request().Context(), contextKey, user)))
}
