package main

import (
	"log"
	"net/http"

	"github.com/azbshiri/hex-rest-api/pkg/http/handler"
)

func main() {
	router := handler.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
