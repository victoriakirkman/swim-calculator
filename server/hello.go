package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	webRootEnv = "SWIM_CALCULATOR_WEBROOT"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
	log.Printf("hello")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	indexLoc := os.Getenv(webRootEnv) + "/index.html"
	fmt.Printf("\nlocation: %v\n", indexLoc)
	fs := http.FileServer(http.Dir(indexLoc))

	log.Printf("Starting server")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error running server: %+v", err)
	}

	http.Handle("/", fs)
}
