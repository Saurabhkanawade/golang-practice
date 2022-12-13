package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Post struct {

	Title  string `json:"title"`

}

func main() {
	fmt.Println("Welcome to the get request")

	response, err := http.Get(`https://jsonplaceholder.typicode.com/posts`)

	newCheckErrors(err)

	log.Print(response.Body)

	getResponse, err := ioutil.ReadAll(response.Body)

	newCheckErrors(err)

	fmt.Println(string(getResponse))

	// unmarshling the json to object

	var post Post

	errUnmarshal,err:= json.Unmarshal(bytes,&post)

	newCheckErrors(err)

	defer response.Body.Close()

}
func newCheckErrors(err error) {
	if err != nil {
		panic(err)
	}
}
