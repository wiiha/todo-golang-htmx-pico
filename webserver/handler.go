package webserver

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type todoItem struct {
	ID   int    `form:"id"`
	What string `form:"what"`
	Done bool   `form:"done"`
}

func (t *todoItem) Valid() bool {
	return t.What != ""
}

func addTodo(c echo.Context) error {

	t := new(todoItem)
	if err := c.Bind(t); err != nil {
		return err
	}

	if !t.Valid() {
		return echo.ErrBadRequest
	}

	log.Printf("adding item: %+v", t)

	return c.Render(http.StatusOK, "todoItem", t)
}
