package main

import (
	"github.com/gkranasinghe/go-api/errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	ServerString        = ":8080"
	ServerTimeoutAmount = 20
)

func NewServer(handler *Handler) (*echo.Echo, error) {
	e := echo.New()
	e.Logger.Print("Starting Main Loop")

	// Skip auth, validation and logging for readiness probe and metrics routes
	skipper := RouteSkipper([]string{"/ready", "/metrics"})

	loggerConfig := middleware.DefaultLoggerConfig
	loggerConfig.Skipper = skipper
	loggerMiddleware := middleware.LoggerWithConfig(loggerConfig)

	e.Use(middleware.Recover())
	e.Use(loggerMiddleware)

	e.HTTPErrorHandler = errors.CustomHTTPErrorHandler

	// RegisterHandlers(e, handler)
	e.POST("/", handler.FindAll)

	return e, nil
}

type Server struct {
	handler *Handler
}

func (s *Server) Run(e *echo.Echo) {
	e.Logger.Fatal(e.Start(":1323"))
}

// type Config struct {
// 	Enabled      bool
// 	DatabasePath string
// 	Port         string
// }

// func NewConfig() *Config {
// 	return &Config{
// 		Enabled:      true,
// 		DatabasePath: "./example.db",
// 		Port:         "8000",
// 	}
// }
