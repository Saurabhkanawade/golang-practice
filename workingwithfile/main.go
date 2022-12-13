package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("The Working with file ")

	content := "Hello my name is saurabh kanawade. I m working as sde in crossayst infotech"

	fileCreate, err := os.Create("./newMyFile.txt")

	checkNilError(err)

	writeData, err := io.WriteString(fileCreate, content)

	fmt.Println("The file is ", writeData)

	defer fileCreate.Close()
}

func readFile(filename string) {
	readData, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("The data of the file is : ", string(readData))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
