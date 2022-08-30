package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Home</h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>About</h1>")
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page not found", http.StatusNotFound)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/about", aboutHandler)
	r.NotFound(errorHandler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
