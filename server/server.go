package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"log"
	"net/http"
	"os"
)

const (
	webRoot = "SWIM_CALCULATOR_WEBROOT"
)

type JsonEncoder interface {
	Encode(v interface{}) ([]byte, error)
	EncodeResponse(status int, v interface{}) error
}

func hello(enc JsonEncoder, params martini.Params) (int, string) {

	fmt.Printf("\nEndpoint requested /hello")
	return http.StatusOK, "hello\n"
}

/* func headers(w http.ResponseWriter, req *http.Request) {

	fmt.Printf("\nEndpoint requested /headers")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}*/

func main() {

	// http.HandleFunc("/hello", hello)
	// http.HandleFunc("/headers", headers)

	// http.ListenAndServe(":8080", nil)

	// m := httputil.SetupMartini(httputil.MartiniLogging os.Getenv(webRoot), "index.html")
	m := martini.New()

	m.Use(func(w http.ResponseWriter) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "SAMEORIGIN")
		w.Header().Set("X-UA-Compatible", "IE=edge")
	})
	m.Use(martini.Static(os.Getenv(webRoot), martini.StaticOptions{IndexFile: "index.html", SkipLogging: true}))

	r := martini.NewRouter()
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)
	classic := &martini.ClassicMartini{
		Martini: m,
		Router:  r,
	}

	classic.Get("/hello", func() string {
		return "hello"
	})
	http.Handle("/", classic)

	log.Printf("Starting server")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error running server: %+v", err)
	}
}
