package webserver

import (
	"log"
	"net/http"

	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo/v4"
)

func Start() {

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
	e.POST("/todo", addTodo)
	e.GET("/", func(c echo.Context) error {
		items := []*todoItem{
			{
				ID:   1,
				What: "check me!",
				Done: false,
			},
			{
				ID:   2,
				What: "and me!",
				Done: false,
			},
			{
				ID:   3,
				What: "I am already done!",
				Done: true,
			},
		}
		return c.Render(http.StatusOK, "index.html", IndexPageData{Todos: items})
	})
	e.Logger.Fatal(e.Start(":1323"))
}

type IndexPageData struct {
	Todos []*todoItem
}
