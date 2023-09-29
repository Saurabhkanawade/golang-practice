package model

type Author struct {
	AuthorId  string `json:"author_id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Phone     string `json:"phone"`
}
