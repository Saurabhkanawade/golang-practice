package repository

import "github.com/volatiletech/null/v8"

type StudentRequest struct {
	Name    null.String `json:"name"`
	College null.String `json:"college"`
	Gender  null.String `json:"gender"`
}

type StudentResponse struct {
	Id int64 `json:"id"`
}

type Student struct {
	Id      int64       `json:"id"`
	Name    null.String `json:"name"`
	College null.String `json:"college"`
	Gender  null.String `json:"gender"`
}
