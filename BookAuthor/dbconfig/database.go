package dbconfig

import (
	"github.com/SaurabhKanawade/BookAuthor/model"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
	"os"
)

func Connect() *pg.DB {

	log.Println("Connecting to the database .................")

	opts := &pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     os.Getenv("DB_HOST"),
		Database: os.Getenv("DB_DATABASE"),
	}

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Println("failed to connect database...........")
		os.Exit(100)
	} else {
		log.Println("Connected success to database ..............")
	}

	// creating the schema if not present
	err := CreateSchema(db)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Disconnected from the database ...............")
	return db

}

//creating table function

func CreateSchema(db *pg.DB) error {

	models := []interface{}{
		(*model.Book)(nil),
		(*model.Author)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})

		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
