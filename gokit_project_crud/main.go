// # Endpoints for Portals
//
// gokit_project_crud:
//
// Schemes: http, https
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package main

import (
	"context"
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/saurabhkanawade/gokit_project_crud/internal/endpoints"
	"github.com/saurabhkanawade/gokit_project_crud/internal/repository"
	"github.com/saurabhkanawade/gokit_project_crud/internal/services"
	"github.com/saurabhkanawade/gokit_project_crud/internal/transport"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

// @title My Go kit Microservice API
// @version 1.0
// @description This is a sample API for a Go kit microservice using Swagger.
// @host localhost:8080
// @BasePath /
func main() {
	logrus.Info("starting gokit crud service")

	dsn := "host=localhost port=5432 user=postgres password=omsairam dbname=gokit sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Test the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	logrus.Info("Connected to the PostgreSQL Database!")
	router := mux.NewRouter().StrictSlash(true)
	log := logrus.New()

	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.TextFormatter{})

	if err != nil {
		logrus.Error("unable to connect to database", err)
	}

	ctx := context.Background()
	var srv services.Service
	{
		repo := repository.NewRespostiry(*db, *log)
		srv = services.NewService(repo, *log)
	}
	endpoint := endpoints.MakeNewEndpoints(srv)

	//// server swagger page
	//fs := http.FileServer(http.Dir("./swagger-ui"))
	//router.PathPrefix("/swagger-ui/").
	//	Handler(http.StripPrefix("/swagger-ui/", fs))
	sh := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./swaggerui/")))
	router.PathPrefix("/swaggerui/").Handler(sh)

	logrus.Info("Starting new server on the  localhost:8080")
	handler := transport.NewHTTPServer(ctx, endpoint)
	logrus.Fatal(http.ListenAndServe(":8080", handler))

}
