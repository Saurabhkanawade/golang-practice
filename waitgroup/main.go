package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	websiteList := []string{
		"https://google.com",
		"https://facebook.com",
		"https://javatpoint.com",
		"https://learnjava.com",
	}
	for _, web := range websiteList {
		go getStatusCode(web)
	}
}

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%d Success the api call status code %s\n", res.StatusCode, endpoint)
}
