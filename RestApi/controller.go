package main

import (
	"github.com/go-pg/pg"
)

func (p *product) GetProducts(db *pg.Conn) {

	query := "select * from product"

	db.Model(&p).Exec(query)
}
