package todo

import (
	"fmt"
	"sort"
	"sync"
	"try-htmx/pkg/utils"
)

type TodoItem struct {
	ID   string `form:"id"`
	What string `form:"what"`
	Done bool   `form:"done"`
}

type TodoItems []*TodoItem

func (t TodoItems) Len() int {
	return len(t)
}

func (t TodoItems) Less(i, j int) bool {
	return t[i].ID < t[j].ID
}

func (t TodoItems) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
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
	ts := TodoItems{}

	for _, v := range svc.todos {
		ts = append(ts, v)
	}

	sort.Sort(ts)

	return ts
}

func (svc *TodoSVC) Get(id string) (*TodoItem, error) {
	svc.RLock()
	defer svc.RUnlock()
	i, ok := svc.todos[id]
	if !ok {
		return nil, fmt.Errorf("no todo with id %s", id)
	}

	return i, nil

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
