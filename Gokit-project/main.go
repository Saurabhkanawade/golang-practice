package main

import (
	"context"
	"fmt"
	_ "github.com/gorilla/mux"
	"github.com/joho/godotenv"
	endpoints2 "github.com/saurabhkanawade/internal/endpoints"
	"github.com/saurabhkanawade/internal/services"
	"github.com/saurabhkanawade/internal/transport"
	"log"
	"net/http"
	"os"
)

func main() {

	fmt.Println("Welcome to go-kit")

	err := godotenv.Load("application.env")
	CheckNilErr(err)
	ctx := context.Background()
	var svc services.Service
	endpoints := endpoints2.MakeEndpoints(svc)
	handler := transport.MakeHTTPHandlers(ctx, endpoints)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func CheckNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
