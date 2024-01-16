package main

import (
	"log"
	"sharing_vasion_indonesia/api_gateway/internal/server"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	s := server.NewServer(r)

	log.Fatal(s.Run())
}
