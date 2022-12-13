package main

import (
	"fmt"
)

type User struct {
	Name   string
	Gmail  string
	Age    int
	Status bool
}

func (u User) getDetails() {
	u.Name = "saurabh kanawade"
	u.Gmail = "saurabhkanawade30@gmail.com"
	u.Age = 50
	u.Status = false
	fmt.Println("The user details in method is : ", u)
}

func main() {
	fmt.Println("Welcome to the methods")

	details := User{"saurabh kanawade", "sk@gmail.com", 22, true}

	fmt.Println("The main method details ", details)

	details.getDetails()

}
