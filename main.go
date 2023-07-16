package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mmiecz/golearn/controllers"
	"github.com/mmiecz/golearn/views"
	"net/http"
	"path/filepath"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Awesome GO!</h1>")
}

func handleContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Info</h1><p>email here</p>")
}
func handleGallery(w http.ResponseWriter, r *http.Request) {
	galleryUid := chi.URLParam(r, "galleryUid")
	fmt.Fprintf(w, "GET gallery %v", galleryUid)
}

func main() {
	r := chi.NewRouter()

	homeTemplate := views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))
	contactTemplate := views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))

	r.Use(middleware.Logger)
	r.Get("/", controllers.StaticHandler(homeTemplate))
	r.Get("/contact", controllers.StaticHandler(contactTemplate))
	r.Get("/gallery/{galleryUid}", handleGallery)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://http.cat/status/404", 404)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe("127.0.0.1:3000", r)
}
