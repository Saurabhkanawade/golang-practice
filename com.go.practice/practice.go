package main

import "fmt"

//globel variable
//var name string = "saurabh"

func myName(str string) string {
	return "Hello " + str + ", welcome to the world of technology ! "
}

func myDreamBike(bikeName string) {
	fmt.Println("My Dream bike is : ", bikeName)
}

func addNumbers(a, b int) int {
	total := a + b
	return total
}

func main() {

	fmt.Println("Hello world")

	fmt.Println(myName("Saurabh Kanawade"))
	myDreamBike("Yamaha M15")
	fmt.Println("The Addition of the two number is ", addNumbers(12, 12))

}
