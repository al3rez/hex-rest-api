package main

import (
	"log"
	"net/http"

	"github.com/azbshiri/hex-rest-api/pkg/auth"
	"github.com/azbshiri/hex-rest-api/pkg/http/handler"
	"github.com/azbshiri/hex-rest-api/pkg/storage/mem"
)

func main() {
	s := new(mem.Storage)
	auth := auth.NewService(s)
	router := handler.NewRouter(auth)
	log.Fatal(http.ListenAndServe(":8080", router))
}
