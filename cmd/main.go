package main

import (
	"log"
	"net/http"

	"github.com/namanag0502/go-blog/pkg/routes"
)

func main() {
	mux := routes.Routes()

	server := &http.Server{
		Addr:    ":4000",
		Handler: mux,
	}

	log.Println("Server is running on port http://localhost:4000/")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
