package main

import (
	"github.com/go-pg/pg"
	"log"
	"os"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *pg.DB
}

func (a *App) Initialize(user, password, dbname string) {

	options := &pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     os.Getenv("DB_ADDR"),
		Database: os.Getenv("DB_DATABASE"),
	}

	a.DB = pg.Connect(options)

	a.Router = mux.NewRouter()

	a.Router.Get("/products")

}

func (a *App) Run(addr string) {}

// error handler

func checkErrorIsNill(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
