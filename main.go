package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	fmt.Println("Server is running on http://localhost:80")
	http.ListenAndServe(":80", nil)
}
