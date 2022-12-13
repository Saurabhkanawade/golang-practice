package main

import "fmt"

func greetings() {
	fmt.Println("Namstey from golang functions")
}

func greetingsTwo() {
	fmt.Println("Namstey India")
}

func adder(value int, valueTwo int) int {
	return value + valueTwo
}

func proAdder(value ...int) (int, string) {
	total := 0

	for _, val := range value {
		total += val
	}
	return total, "The value are "
}

func main() {

	fmt.Println("Welcome to the function in golang")

	greetingsTwo()
	greetings()

	result := adder(3, 3)

	fmt.Println(result)

}
