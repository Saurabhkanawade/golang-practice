package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the RestApi")

	a := App{}

	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":9090")
}
