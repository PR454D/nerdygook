package main

import (
	"log"
	"net/http"

	"github.com/PR454D/nerdygook/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterNerdyGookRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9000", r))
}
