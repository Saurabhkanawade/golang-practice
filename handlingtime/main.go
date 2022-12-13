package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Handling time in golang")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))

	createDate := time.Date(2022, time.September, 30, 12, 00, 00, 00, time.UTC)
	fmt.Println(createDate)

	fmt.Println(createDate.Format("02-01-2006 Monday"))

}
