package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mmiecz/golearn/controllers"
	"github.com/mmiecz/golearn/templates"
	"github.com/mmiecz/golearn/views"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	homeTemplate := views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))
	contactTemplate := views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))

	var usersController controllers.Users
	usersController.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Use(middleware.Logger)
	r.Get("/", controllers.StaticHandler(homeTemplate))
	r.Get("/contact", controllers.StaticHandler(contactTemplate))
	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))
	r.Get("/signup", usersController.New)
	r.Post("/signup", usersController.Create)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://http.cat/status/404", 404)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe("127.0.0.1:3000", r)
}
