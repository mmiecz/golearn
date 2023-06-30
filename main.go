package main

import (
	"fmt"
	"net/http"
)

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		handleHome(w, r)
	case "/contact":
		handleContact(w, r)
	default:
		handleNotFound(w, r)
	}
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Page Not Found!")
}
func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Awesome GO!</h1>")
}

func handleContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Info</h1><p>email here</p>")
}

func main() {
	var router Router
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe("127.0.0.1:3000", router)
}
