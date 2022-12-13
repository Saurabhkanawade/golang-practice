package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Lets do the code of conversion")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for app")

	input, _ := reader.ReadString('\n')

	fmt.Println("The rate of out app is : ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your application ", numRating+1)
		fmt.Println()
	}

}
