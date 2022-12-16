package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {

	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	checkErrorIsNill(err)

	a.Router = mux.NewRouter()

}

func (a *App) Run(addr string) {}

// error handler

func checkErrorIsNill(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
