package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	ignore := false
	if r.Method != http.MethodGet {
		// ignore if anything but a GET is sent
		ignore = true
	}
	if r.URL.Path != "/" {
		// ignore if anything but the correct path is sent
		ignore = true
	}
	if !ignore {
		fmt.Fprint(w, "hello world")
	}

}

func main() {
	http.HandleFunc("/", HelloWorldHandler)

	log.Fatal(http.ListenAndServe(":3333", nil))

}
