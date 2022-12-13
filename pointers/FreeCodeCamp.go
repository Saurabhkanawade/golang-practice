package main

import "fmt"

func changeValue(string2 *string) {
	*string2 = "new change value"
}

func main() {

	myNumber := 50

	myString := "saurabh kanawade"

	var newNumber *int = &myNumber // & oprator taking the address location of the my number

	fmt.Println(myNumber)
	fmt.Println(*newNumber) // Dereferencing a pointer gives us access to the value the pointer points to.

	var thirdNumber **int = &newNumber

	fmt.Println(**thirdNumber)

	changeValue(&myString)
	fmt.Println(myString)

	fmt.Println("*********************")

	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

}
