package main

import "fmt"

func main() {

	fmt.Println("WElcome the loops")

	days := []string{"Sun", "Mon", "Tues", "Wens", "Thir", "FRi"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	fmt.Println()

	for i := range days {
		fmt.Println(days[i])
	}
	fmt.Println()

	for index, value := range days {
		fmt.Printf("The index %v has value %v\n :", index, value)
	}

	fmt.Println("Goto , break and Continue statements")

}
