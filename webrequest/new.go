package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.youtube.com/watch?v=ru53LpdVHn4&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=25"

func main() {
	fmt.Print("Welcome to the new go file")

	response, err := http.Get(url)
	newCheckError(err)

	getData, err := ioutil.ReadAll(response.Body)
	newCheckError(err)

	content := string(getData)

	fmt.Println(content)

	defer response.Body.Close()

}

func newCheckError(err error) {
	if err != nil {
		panic(err)
	}
}
