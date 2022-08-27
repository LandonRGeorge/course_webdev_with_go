package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "This is my home page")
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<strong>About</strong> page")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	log.Println(http.ListenAndServe(":3000", nil))
}
