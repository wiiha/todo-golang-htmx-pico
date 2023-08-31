package todo

import (
	"testing"
)

func TestAdd(t *testing.T) {
	what := "my test item"
	svc := NewTodoSVC()

	id, err := svc.Add(what)
	if err != nil {
		t.Fatalf("adding item: %v", err)
	}
	if id == "" {
		t.Fatalf("expected id != \"\"")
	}
}

func TestList(t *testing.T) {
	ts := []*TodoItem{}

	ts = append(ts, NewTodoItem("1"))
	ts = append(ts, NewTodoItem("2"))
	ts = append(ts, NewTodoItem("3"))

	svc := NewTodoSVC()

	for _, t := range ts {
		svc.todos[t.ID] = t
	}

	items := svc.List()

	if len(items) != 3 {
		t.Fatalf("Expected 3 items, got %d", len(items))
	}
}

func TestMarkAsDone(t *testing.T) {
	ts := []*TodoItem{}

	item := NewTodoItem("my test")

	ts = append(ts, item)

	svc := NewTodoSVC()

	for _, t := range ts {
		svc.todos[t.ID] = t
	}

	if err := svc.MarkAsDone(item.ID); err != nil {
		t.Fatalf("marking as done: %v", err)
	}

	if !item.Done {
		t.Fatal("expected item to be marked as done.")
	}
}
