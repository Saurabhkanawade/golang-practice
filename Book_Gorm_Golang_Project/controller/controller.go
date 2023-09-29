package controller

import (
	"github.com/gofiber/fiber"
	"github.com/saurabhkanawade/book_gorm_golang_project/database"
	"github.com/saurabhkanawade/book_gorm_golang_project/helper"
	"github.com/saurabhkanawade/book_gorm_golang_project/model"
	"log"
)

func GetBooks(ctx *fiber.Ctx) {
	log.Println("fetching all books from the database.................")

	var books []model.Book
	database.DB.Model(&model.Book{}).Preload("Author").Find(&books)

	ctx.JSON(books)
	log.Println("fetch success books.................")
}

func GetBook(ctx *fiber.Ctx) {
	log.Println("fetching the book ..........")

	book := model.Book{}
	bookId := ctx.Params("id")
	database.DB.Model(&model.Book{}).Preload("Author").Find(&book, bookId)
	ctx.JSON(book)

	log.Println("fetched success book with BookId :", bookId)
}

func CreateBooks(ctx *fiber.Ctx) {
	log.Println("Creating new book .................")

	book := model.Book{}
	err := ctx.BodyParser(&book)
	helper.HandleNilError(err)

	result := database.DB.Create(&book)
	errs := ctx.JSON(book)
	helper.HandleNilError(errs)

	log.Println("created new book success ....", result)
}

func UpdateBook(ctx *fiber.Ctx) {
	log.Println("updating the existing book..........")

	book := model.Book{}
	author := model.Author{}
	bookId := ctx.Params("id")

	database.DB.Find(&book, bookId)
	err := ctx.BodyParser(&book)
	helper.HandleNilError(err)

	database.DB.Find(&author, &book.AuthorId)
	book.AuthorId = author.ID
	book.Author.Name = author.Name
	book.Author.Address = author.Address

	database.DB.Save(&book)
	database.DB.Save(&author)
	errs := ctx.JSON(book)
	helper.HandleNilError(errs)

	log.Println("updated success the book with id..", book.ID)
}

func DeleteBook(ctx *fiber.Ctx) {
	log.Println("deleting the book from the database .........")

	book := model.Book{}
	bookId := ctx.Params("id")
	database.DB.Delete(&book, bookId)
	database.DB.Delete(&book, book.AuthorId)
	ctx.JSON(book)

	log.Println("deleted success book from the database .........")
}
