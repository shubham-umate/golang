package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang, no super, no parent

	shubham := User{"Shubham", "shubham@go.dev", true, 19}
	fmt.Println(shubham)
	fmt.Printf("shubham details are %+v\n", shubham)
	fmt.Printf("shubham Name is %v\n", shubham.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
