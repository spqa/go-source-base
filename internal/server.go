package internal

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
	"mcm-api/config"
	_ "mcm-api/docs"
	"mcm-api/internal/router"
	"mcm-api/pkg/log"
	"mcm-api/pkg/response"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	config         *config.Config
	echo           *echo.Echo
	userRouter     *router.UserRouter
	documentRouter *router.DocumentRouter
}

func newServer(
	config *config.Config,
	userRouter *router.UserRouter,
	documentRouter *router.DocumentRouter,
) *Server {
	e := echo.New()
	e.HideBanner = true
	e.HTTPErrorHandler = errorHandler
	return &Server{
		config:         config,
		echo:           e,
		userRouter:     userRouter,
		documentRouter: documentRouter,
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
func (s Server) registerRouter() {
	s.echo.GET("/swagger/*", echoSwagger.WrapHandler)
	s.userRouter.Register(s.echo.Group("users"))
	s.documentRouter.Register(s.echo.Group("documents"))
}

func (s Server) registerMiddleware() {
	s.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:4200",
			"https://localhost:4200",
			s.config.WebAppUrl,
		},
	}))
}

func errorHandler(err error, c echo.Context) {
	if !c.Response().Committed {
		switch e := err.(type) {
		case *response.ApiError:
			_ = c.JSON(e.StatusCode, e)
			break
		default:
			log.Logger.Error("unhandled error")
			internalError := response.NewApiInternalError(nil)
			_ = c.JSON(internalError.StatusCode, internalError)
		}
	}
}

func (s *Server) Start() {
	s.registerRouter()
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
