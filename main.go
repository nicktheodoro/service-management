package main

import (
	"net/http"

	"github.com/go-playground/pure/v5"
	mw "github.com/go-playground/pure/v5/_examples/middleware/logging-recovery"
)

func main() {
	p := pure.New()
	p.Use(mw.LoggingAndRecovery(true))

	p.Get("/", helloWorld)
	p.Get("/users", helloUser)

	http.ListenAndServe(":3007", p.Serve())
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func helloUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello User"))
}
