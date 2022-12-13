package main

import "fmt"

func main() {

	fmt.Println("Welcome to the ifelse")

	if num := 5; num <= 5 {
		fmt.Println("The no is is less than 5")
	} else {
		fmt.Println("The no is greater than 5")
	}

	if age := 18; age < 18 {
		fmt.Println("You can not drive bike or car")
	} else {
		fmt.Println("You can drive car and bike")
	}

	loginCount := 15

	var result string

	if loginCount <= 10 {
		result = "Regular user"
	} else if loginCount >= 10 && loginCount <= 15 {
		result = "New User"
	} else {
		fmt.Println("!15 to 20 people")
	}

	fmt.Println(result)

	numCheck := 11

	if numCheck%2 == 0 {
		fmt.Println("The no is even")
	} else {
		fmt.Println("The no is odd")
	}

	numArray := []int{1, 2, 3, 4, 5, 6, 7}

	for index, value := range numArray {

		if value%2 == 0 {
			fmt.Printf("even no with index %v and the value is %v\n:", index, value)
		} else {
			fmt.Printf("odd no with index %v and the value is %v \n", index, value)
		}
	}

}
