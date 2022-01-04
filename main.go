package main

import (
	"book-server/config"
	"book-server/router"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := start(); err != nil {
		log.Fatal(err)
	}

}

func start() error {
	ctx := context.Background()
	cfg := config.Get()

	bookRouter := router.NewBookRouter(ctx)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group("/v1")
	bookRoutes := v1.Group("/book")
	// bookRoutes.GET("/", bookRouter.GetAll)
	// bookRoutes.POST("/", bookRouter.Create)
	bookRoutes.GET("/:id", bookRouter.Get)
	// bookRoutes.PUT("/:id", bookRouter.Update)
	// bookRoutes.DELETE("/:id", bookRouter.Delete)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})

	s := &http.Server{
		Addr:         cfg.HTTPAddr,
		ReadTimeout:  30 * time.Minute,
		WriteTimeout: 30 * time.Minute,
	}

	e.Logger.Fatal(e.StartServer(s))

	return nil
}
