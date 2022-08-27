package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Home</h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>About</h1>")
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Couldn't find that page!")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		aboutHandler(w, r)
	default:
		errorHandler(w, r)
	}
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		aboutHandler(w, r)
	default:
		errorHandler(w, r)
	}
}

func main() {
	var router Router
	//http.HandleFunc("/", pathHandler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", router)
}
