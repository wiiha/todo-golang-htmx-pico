package webserver

import (
	"html/template"
	"io"
	"strings"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplater(pathTmplDir string) *Template {
	p := strings.TrimSuffix(pathTmplDir, "/")
	return &Template{
		templates: template.Must(template.ParseGlob(p + "/*.html")),
	}
}
