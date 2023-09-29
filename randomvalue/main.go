package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("The random value and Math crypto package in golang")

	rand.Seed(time.Now().UnixNano())

	//for i := 1; i < 40; i++ {
	//	fmt.Println(rand.Intn(39))
	//}

}
