package database

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/saurabhkanawade/customerprojectpg_fb/model"
	"log"
	"os"
)

func Connect() *pg.DB {
	options := &pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     os.Getenv("DB_ADDR"),
		Database: os.Getenv("DB_DATABASE"),
	}

	var db *pg.DB = pg.Connect(options)

	if db == nil {
		log.Printf("Failed to connect ...........\n")
		os.Exit(100)
	}
	log.Printf("Connected sucesss...\n")

	err := createSchema(db)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func createSchema(db *pg.DB) interface{} {

	models := []interface{}{
		(*model.Customer)(nil),
	}

	for _, mod := range models {
		err := db.Model(mod).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
