package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "What is going on there?")
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil)
}
