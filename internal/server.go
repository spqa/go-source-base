package internal

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
	"mcm-api/config"
	_ "mcm-api/docs"
	"mcm-api/pkg/document"
	"mcm-api/pkg/log"
	"mcm-api/pkg/user"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	config          *config.Config
	echo            *echo.Echo
	userHandler     *user.Handler
	documentHandler *document.Handler
}

func newServer(
	config *config.Config,
	userRouter *user.Handler,
	documentRouter *document.Handler,
) *Server {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:4200",
			"https://localhost:4200",
			config.WebAppUrl,
		},
	}))
	return &Server{
		config:          config,
		echo:            e,
		userHandler:     userRouter,
		documentHandler: documentRouter,
	}
}

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
func (s Server) registerHandler() {
	s.echo.GET("/swagger/*", echoSwagger.WrapHandler)
	s.userHandler.Register(s.echo.Group("users"))
	s.documentHandler.Register(s.echo.Group("documents"))
}

func (s *Server) Start() {
	s.registerHandler()
	go func() {
		if err := s.echo.Start(":3000"); err != nil {
			log.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.echo.Shutdown(ctx); err != nil {
		log.Logger.Fatal("Error shutting down server", zap.Error(err))
	}
}
