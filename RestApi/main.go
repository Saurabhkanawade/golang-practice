package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the RestApi")

	a := App{}

	a.Initialize(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"))

	a.Run(":9090")
}
