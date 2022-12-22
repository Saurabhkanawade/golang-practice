package main

import (
	"fmt"
	"time"
)

func main() {

	greeter("Hello")
	greeter("World")
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println(s)
	}
}
