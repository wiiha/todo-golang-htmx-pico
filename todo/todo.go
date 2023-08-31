package todo

import (
	"fmt"
	"sync"
	"try-htmx/pkg/utils"
)

type TodoItem struct {
	ID   string `form:"id"`
	What string `form:"what"`
	Done bool   `form:"done"`
}

func (t *TodoItem) Valid() bool {
	return t.What != ""
}

func NewTodoItem(what string) *TodoItem {
	return &TodoItem{
		ID:   utils.MustNewID(),
		What: what,
		Done: false,
	}
}

type TodoSVC struct {
	todos map[string]*TodoItem
	*sync.RWMutex
}

func NewTodoSVC() *TodoSVC {
	return &TodoSVC{
		todos:   make(map[string]*TodoItem),
		RWMutex: new(sync.RWMutex),
	}
}

func (svc *TodoSVC) List() []*TodoItem {
	svc.RLock()
	defer svc.RUnlock()
	ts := []*TodoItem{}

	for _, v := range svc.todos {
		ts = append(ts, v)
	}

	return ts
}

func (svc *TodoSVC) add(t *TodoItem) (string, error) {
	if !t.Valid() {
		return "", fmt.Errorf("todo item is not valid")
	}
	svc.Lock()
	defer svc.Unlock()

	svc.todos[t.ID] = t

	return t.ID, nil

}

func (svc *TodoSVC) Add(what string) (string, error) {

	return svc.add(NewTodoItem(what))

}

func (svc *TodoSVC) MarkAsDone(id string) error {
	svc.Lock()
	defer svc.Unlock()

	t, ok := svc.todos[id]

	if !ok {
		return fmt.Errorf("no item for id %s", id)
	}

	t.Done = true

	return nil

}
