package main

import "fmt"

func main() {

	map1 := map[int]string{
		1: "saurabh",
		2: "aalam",
	}

	fmt.Println("The list is ", map1)

	programmingLanguages := make(map[int]string)

	programmingLanguages[1] = "Java"
	programmingLanguages[2] = "Python"
	programmingLanguages[3] = "Golang"
	programmingLanguages[4] = "Ruby"

	//fmt.Println(programmingLanguages)
	//fmt.Println(programmingLanguages[2])
	//
	//delete(programmingLanguages, 1)
	//fmt.Println(programmingLanguages)

	//loop

	for key, values := range programmingLanguages {
		fmt.Println(key, values)

		fmt.Printf("The key %v has value : %v", key, values)
	}

	for key, values := range programmingLanguages {
		fmt.Println(values)

		fmt.Printf("The key %v has value : %v", key, values)
	}

}
