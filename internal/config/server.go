package config

import (
	"net/http"

	"github.com/krishnapramodaradhi/xpressbuy-api/internal/entity"
	"github.com/krishnapramodaradhi/xpressbuy-api/internal/handler"
	m "github.com/krishnapramodaradhi/xpressbuy-api/internal/middleware"
	customerror "github.com/krishnapramodaradhi/xpressbuy-api/internal/util/customError"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
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
	app.Logger.SetLevel(log.Lvl(1))
	app.Validator = m.New()
	app.HTTPErrorHandler = new(customerror.CustomError).ErrorHandler

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

	// Protected Routes
	// Cart Routes
	p := app.Group("/api/v1/cart")
	p.Use(m.ValidateToken)
	ch := handler.NewCartHandler(db.db)
	p.GET("", ch.FetchCart)
	p.POST("/modify", ch.AddItemToCart)
	p.DELETE("/remove/:id", ch.RemoveFromCart)
	p.DELETE("/remove", ch.ClearCart)

	app.Logger.Fatal(app.Start(s.listenAddr))
}
