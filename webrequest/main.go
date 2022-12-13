package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Print("Welcome to the webrequest")

	const url = "https://go.dev/"

	response, err := http.Get(url)
	checkNilError(err)

	fmt.Printf("The response of the http request %T:", response)

	data, err := ioutil.ReadAll(response.Body)

	checkNilError(err)

	content := string(data)
	fmt.Println(content)

	defer response.Body.Close() //close the connection in the last

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
