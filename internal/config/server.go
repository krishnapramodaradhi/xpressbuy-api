package config

import (
	"net/http"

	"github.com/krishnapramodaradhi/xpressbuy-api/internal/entity"
	"github.com/krishnapramodaradhi/xpressbuy-api/internal/handler"
	m "github.com/krishnapramodaradhi/xpressbuy-api/internal/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{listenAddr: listenAddr}
}

func (s *Server) Run() {
	app := echo.New()

	app.HideBanner = true

	// middlewares
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	// DB initialization
	db, err := NewDatabase()
	if err != nil {
		app.Logger.Fatal("there was an error while connecting to DB", err)
	}

	// Routes
	app.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, entity.NewHealthCheck())
	})

	// Product Routes
	r := app.Group("/api/v1/products")
	ph := handler.NewProductHandler(db.db)
	r.GET("", ph.FetchProducts)
	r.GET("/:id", ph.FetchProductById)

	// Auth Routes
	r = app.Group("/api/v1/auth")
	ah := handler.NewAuthHandler(db.db)
	r.POST("/signup", ah.Register)
	r.POST("/signin", ah.Login)

	// Protected Route Group
	p := app.Group("/api/v1/cart")
	ch := handler.NewCartHandler(db.db)
	p.POST("/add", ch.AddItemToCart, m.ValidateToken)

	app.Logger.Fatal(app.Start(s.listenAddr))
}
