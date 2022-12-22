package main

import (
	"github.com/saurabhkanawade/router"
	"log"
	"net/http"
)

func main() {
	log.Println("The server is getting started on port:9000........")

	r := router.Router()

	http.ListenAndServe(":9000", r)
	//server conf
	log.Fatal(http.ListenAndServe(":9000", r))

}
