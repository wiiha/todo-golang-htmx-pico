package webserver

import (
	"log"
	"try-htmx/todo"

	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo/v4"
)

func Start() {

	todoSvc := todo.NewTodoSVC()
	handlers := NewHandlerSvc(todoSvc)

	t := NewTemplater("webserver/views")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	e := echo.New()
	e.Renderer = t

	// Assign new renderer on file change.
	// Avoids need for server restart.
	go func(e *echo.Echo) {
		for {
			select {
			case event := <-watcher.Events:
				// log.Println("Event:", event)
				// Check if the event is a write event
				if event.Has(fsnotify.Chmod) {
					continue
				}
				log.Println("Reloading renderer")
				t := NewTemplater("webserver/views")
				e.Renderer = t
			case err := <-watcher.Errors:
				log.Println("Error:", err)
			}
		}
	}(e)

	err = watcher.Add("webserver/views")
	if err != nil {
		log.Fatal(err)
	}

	e.Static("/static", "webserver/static")
	e.POST("/todo", handlers.addTodo)
	e.GET("/", handlers.listTodos)
	e.Logger.Fatal(e.Start(":1323"))
}

type IndexPageData struct {
	Todos []*todo.TodoItem
}
