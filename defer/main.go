package main

import "fmt"

func main() {

	fmt.Println("The defer in go ")

	defer fmt.Println("Hello this is one")
	defer fmt.Println("Hello this is two")
	defer fmt.Println("Hello this is three")
	defer fmt.Println("Hello this is four")

	for i := 0; i < 5; i++ {
		fmt.Println("Welcome to the golang", i)
	}
}
