package main

import (
	"github.com/joho/godotenv"
	"github.com/saurabhkanawade/customerprojectpg_fb/router"
)

func main() {

	godotenv.Load("app.env")

	app := router.Router()

	app.Listen("9000")
}
