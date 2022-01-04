package router

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookRouter struct {
	ctx context.Context
}

func NewBookRouter(ctx context.Context) *BookRouter {
	return &BookRouter{
		ctx: ctx,
	}
}

func (r *BookRouter) GetAll(c echo.Context) error {
	return c.String(http.StatusOK, "get books call")
}

func (r *BookRouter) Get(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "get book : "+id)
}
