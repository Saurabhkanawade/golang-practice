package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("welcome to the go")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please rate our website")

	input, err := reader.ReadString('\n')

	fmt.Println("The rating is :", input)
	fmt.Printf("Thanks for giving rate : %T \n", input, err)

	//
	//fmt.Println("Topic Comma Ok syntax and working in go")
	//
	////creating new reader
	//reader := bufio.NewReader(os.Stdin)
	//
	//fmt.Println("Enter the rating for our pizza : ")
	//
	//// taking the input from user
	//input, _ := reader.ReadString('\n')
	//
	//fmt.Println("Thanks for the rating : ", input)
	//fmt.Printf("Thanks for the rating %T : ", input)

	//comma ok syntax && use _ if you don't care about that variable

	//fmt.Println("Welcome to the CommaOk syntax")
	//
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Println("Enter the rating for out pizza :")
	//
	//input, _ := reader.ReadString('\n')
	//fmt.Println("Thanks for rating :", input)
	//
	//fmt.Printf("Thanks for this rating %T ", input)

	//fmt.Println("Welcome to the CommaOk syntax")
	//
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Println("Enter the rating for out pizza :")
	//
	//input, _ := reader.ReadString('\n')
	//fmt.Println("Thanks for rating :", input)

	//numRating, err := strconv.ParseFloat(input, 64)
	//if  err ! = nil {
	//	fmt.Println(err),
	//}
	//else {
	//	fmt.Println(" Added 1 to your rating : ", numRating+1)
	//}

}
