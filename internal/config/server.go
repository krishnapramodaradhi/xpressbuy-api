package config

import (
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	listenAddr string
}

func (s *Server) Run() {
	app := echo.New()

	// middlewares
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	// Routes
	app.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"id": strconv.Itoa(os.Getpid()), "environment": os.Getenv("APP_ENV"), "status": "healthy"})
	})

	app.Logger.Fatal(app.Start(s.listenAddr))
}

func NewServer(listenAddr string) *Server {
	return &Server{listenAddr: listenAddr}
}
