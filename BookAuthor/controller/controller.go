package controller

import (
	"encoding/json"
	"github.com/SaurabhKanawade/BookAuthor/dbconfig"
	"github.com/SaurabhKanawade/BookAuthor/model"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating new book ..............")
	w.Header().Set("Content-Type", "application/json")

	//connect db
	db := dbconfig.Connect()
	defer db.Close()

	// instance
	var book = &model.Book{
		BookId: uuid.New().String(),
	}

	_ = json.NewDecoder(r.Body).Decode(&book)

	//creating new book

	_, err := db.Model(book).Insert()
	CheckNilError(err)

	//getting response

	_ = json.NewEncoder(w).Encode(book)

	log.Println("Created success new book in databse...........")
}

func CheckNilError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
