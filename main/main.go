package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// função principal
func main() {
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
