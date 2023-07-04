package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	htmlTemplate *template.Template
}

func Parse(filepath string) (Template, error) {
	htmlTemplate, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template %v", err)
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
