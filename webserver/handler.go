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

func (h *handlerSvc) getTodo(c echo.Context) error {
	tid := c.Param("id")

	t, err := h.todoSvc.Get(tid)
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.Render(http.StatusOK, "todoItem", t)
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

type UpdateRequest struct {
	What string `form:"what"`
}

func (h *handlerSvc) updateTodo(c echo.Context) error {

	r := new(UpdateRequest)

	if err := c.Bind(r); err != nil {
		return err
	}
	if r.What == "" {
		return echo.ErrBadRequest
	}

	tid := c.Param("id")

	t := todo.NewTodoItem(r.What)
	t.ID = tid

	if err := h.todoSvc.Update(t); err != nil {
		c.Logger().Warnf("updating item: %v", err)
		return echo.ErrBadRequest
	}

	return c.Render(http.StatusOK, "todoItem", t)
}

func (h *handlerSvc) todoDone(c echo.Context) error {
	tid := c.Param("id")

	if err := h.todoSvc.MarkAsDone(tid); err != nil {
		return echo.ErrBadRequest
	}

	return h.renderTodoList(c)
}

func (h *handlerSvc) renderTodoList(c echo.Context) error {
	ts := h.todoSvc.List()

	return c.Render(http.StatusOK, "todoList", IndexPageData{Todos: ts})
}

func (h *handlerSvc) editTodo(c echo.Context) error {
	tid := c.Param("id")

	t, err := h.todoSvc.Get(tid)
	if err != nil {
		return echo.ErrBadRequest
	}

	return c.Render(http.StatusOK, "editTodo", t)
}

func (h *handlerSvc) deleteTodo(c echo.Context) error {
	tid := c.Param("id")

	if err := h.todoSvc.Delete(tid); err != nil {
		return echo.ErrBadRequest
	}

	return c.NoContent(http.StatusOK)
}
