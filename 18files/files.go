package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Lets learn Files in golang")
	content := "This needs to go in file"

	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Lenght", length)
	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(fileName string)  {
	databyte,err := ioutil.ReadFile(fileName)
	if err!= nil {
		panic(err)
	}

	fmt.Println("Text data inside a file is\n", string(databyte))
}
