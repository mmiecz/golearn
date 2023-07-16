package controllers

import (
	"github.com/mmiecz/golearn/views"
	"html/template"
	"net/http"
)

type Static struct {
	Template views.Template
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML // Controlled by us, no worry about XSS
	}{
		{
			Question: "Question 1?",
			Answer:   "Answer 1",
		},
		{
			Question: "Question 2?",
			Answer:   "Placeholder 2",
		},
		{
			Question: "Question 3?",
			Answer:   "Placeholder 3",
		},
		{
			Question: "Question 4?",
			Answer:   "Placeholder 4",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}

func (static Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	static.Template.Execute(w, nil)
}

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
