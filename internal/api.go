package internal

import (
	"github.com/labstack/echo/v4"
)
import "github.com/swaggo/echo-swagger"

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email superquanganh@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func NewApi() {
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	CreateUserRouter().Register(e.Group("user"))
	e.Logger.Fatal(e.Start(":3000"))
}
