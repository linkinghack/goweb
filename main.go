package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Req: %s \n", r.URL.Path)
		fmt.Fprintf(w, "Hello, Go web. requested: %s \n", r.URL.Path)
	})

	http.ListenAndServe(":8088", nil)
}
