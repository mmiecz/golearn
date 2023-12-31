package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

type Template struct {
	htmlTemplate *template.Template
}

func ParseFS(fs fs.FS, pattern ...string) (Template, error) {
	htmlTemplate, err := template.ParseFS(fs, pattern...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template %w", err)
	}
	return Template{
		htmlTemplate: htmlTemplate,
	}, nil
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTemplate.Execute(w, data)
	if err != nil {
		log.Printf("error while executing template %v", err)
		http.Error(w, "There was an error while generating the template", http.StatusInternalServerError)
		return
	}
}
