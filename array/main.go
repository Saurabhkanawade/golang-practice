package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the array .......")

	//assign array
	var myFruits = []string{"apple", "banana"}

	//myFruits[0] = "Apple"
	//myFruits[1] = "banana"
	//myFruits[2] = "cherry"
	//myFruits[4] = "akhrod"

	//inline declaration + assign
	var myArray = [5]string{"saurabh", "shailesh", "rushikesh", "sagar", "shankar"}

	var bikeList = [2]string{"Yamaha MT 15", " Suziki Baleno "}

	//undefind size array
	var cityName = []string{"sangamner", "pune", "navi mumbai", "mumbai"}

	fmt.Println(myArray)
	fmt.Println("The length Of the Array is : ", len(myArray))
	//[saurabh shailesh rushikesh sagar shankar]

	fmt.Println(myFruits)
	//[Apple banana cherry  akhrod]
	// it gives empty space because of 3 index pos not present

	fmt.Println("The list of bike : ", bikeList)

	fmt.Println(cityName)
	fmt.Println("The Undefind size array length :", len(cityName))

	// adding the values in array if the size is not declare
	myFruits = append(myFruits, "stoberry", "mango", "almond")
	fmt.Println(myFruits)

	fmt.Printf("%T", myFruits)
	var num = 2
	fmt.Println(cityName[:num], cityName[num+1:])
	cityName = append(cityName[:num], cityName[num+1:]...)

	fmt.Println(cityName)

}
