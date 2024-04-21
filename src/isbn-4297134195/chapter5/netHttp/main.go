package main

import (
	"fmt"

	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Fprintf(w, "Hello, World! (GET)")
		case http.MethodPost:
			fmt.Fprintf(w, "Hello, World! (POST)")
		default:
		}
	})

	http.ListenAndServe(":8080", nil)
}
