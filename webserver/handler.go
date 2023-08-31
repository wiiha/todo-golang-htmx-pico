package webserver

import (
	"net/http"
	"try-htmx/todo"

	"github.com/labstack/echo/v4"
)

type handlerSvc struct {
	todoSvc *todo.TodoSVC
}

func NewHandlerSvc(todoSvc *todo.TodoSVC) *handlerSvc {
	return &handlerSvc{
		todoSvc: todoSvc,
	}
}

func (h *handlerSvc) listTodos(c echo.Context) error {
	ts := h.todoSvc.List()
	return c.Render(http.StatusOK, "index.html", IndexPageData{Todos: ts})
}

func (h *handlerSvc) addTodo(c echo.Context) error {

	t := new(todo.TodoItem)
	if err := c.Bind(t); err != nil {
		return err
	}

	if !t.Valid() {
		return echo.ErrBadRequest
	}

	// We only care about "what" prop

	id, err := h.todoSvc.Add(t.What)
	if err != nil {
		c.Logger().Errorf("adding item: %v", err)
		return echo.ErrInternalServerError
	}

	t.ID = id

	return c.Render(http.StatusOK, "todoItem", t)
}
