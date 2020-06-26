package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go is up and running!")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":80", nil)
}
