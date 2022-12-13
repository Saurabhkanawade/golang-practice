package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Welcome to dice ")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	fmt.Printf("Value of dice is %v \n:", diceNumber)

	switch diceNumber {

	case 1:
		fmt.Println("Dice value is 1 and you can open")

	case 2:
		fmt.Println("Dice value is 2 move 2 spots")

	case 3:
		fmt.Println("Dice value is 3 move 3 spots")

	case 4:
		fmt.Println("Dice value is 4 move 4 spots")
		fallthrough
	case 5:
		fmt.Println("Dice value is 5 move 5 spots")

	case 6:
		fmt.Println("Dice value is 6 move 6 spots")

	default:
		fmt.Println("Please enter value between 1 to 6")
	}
}
