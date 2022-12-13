package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
	Gmail     string
	Address   string
}

type Customer struct {
	CustomerId      int
	CustomerName    string
	CustomerAddress string
}

func main() {
	fmt.Println("Welcome to struct in golang")

	data := User{"saurabh", "kanawade", "saurabhkanawade30@gmail.com", "sangamner"}

	fmt.Println("The User struct is  : ", data.FirstName, " | ", data.LastName, " | ", data.Gmail, " | ", data.Address)
	fmt.Printf("\n The User Struct is : \n FirstName : %v\n LastName: %v\n Gmail : %v\n Address : %v \n", data.FirstName, data.LastName, data.Gmail, data.Address)

}
