package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Req: %s \n", r.URL.Path)
		fmt.Fprintf(w, "Hello, this is Go web! requested url: %s, Request time: %+v", r.URL.Path, time.Now())
	})

	fmt.Println("Start listening on port 8088..")
	http.ListenAndServe("0.0.0.0:8088", nil)
}
