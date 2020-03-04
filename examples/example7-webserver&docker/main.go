package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "this is a test server")
	})
	http.ListenAndServe(":"+PORT, nil)
}
