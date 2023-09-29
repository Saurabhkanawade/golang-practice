package model

type Book struct {
	BookId string  `json:"book_id"`
	Title  string  `json:"title"`
	Sells  string  `json:"sells"`
	Author *Author `json:"author"`
}
