package main

import "fmt"

const loginToken string ="jdgsdgjlksjdglkjs"

func main() {

	var username string = "saurabh kanawade"

	fmt.Println("The name of the user is ", username)
	fmt.Printf("Variable is of type : %T \n", username)
	
	var newString string
	fmt.Println(newString)
	
	var website ="www.google.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type : %T \n", website)
	
	
	numberOfUSer := 42.52
	fmt.Println(numberOfUSer)
	fmt.Printf("Variable is of type  : %T \n", numberOfUSer)
	
	
	fmt.Println(loginToken)
	fmt.Printf("Variable is of type  : %T \n", loginToken)
}
