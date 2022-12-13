package main

import "fmt"

func main() {

	fmt.Println("Welcome to slice")

	mySlice := []int{1, 2, 4, 56}
	fmt.Println(mySlice)

	//mySlice=append()

	// adding the values in array if the size is not declare

	myFruits := []string{"banana"}
	myFruits = append(myFruits, "stoberry", "mango", "almond")
	fmt.Println(myFruits)

	myFruits = append(myFruits[:2], myFruits[3:]...)
	fmt.Println(myFruits)

}
