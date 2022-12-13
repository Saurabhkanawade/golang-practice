package main

import "fmt"

func main() {

	fmt.Println("Lets see the workflow the pointers")

	myNumber := 20
	var newPointer = &myNumber

	var newAddress string = "Sangamner"
	var pointAddress *string = &newAddress
	var newPointAddress **string = &pointAddress

	fmt.Println("Value of actual pointer is : ", myNumber)
	fmt.Println("Value of actual pointer is : ", newPointer)
	fmt.Println("Value of actual pointer is : ", *newPointer)
	fmt.Println("The value of new address :", newAddress)
	fmt.Println("The value of new address :", *pointAddress)

	fmt.Println("The value of newPointAddress :", newPointAddress)
	fmt.Println("The value of newPointAddress :", **newPointAddress)

}
