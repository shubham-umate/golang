package main

import "fmt"

func main() {
	fmt.Println("Methods in golang")
	// no inheritance in golang, no super, no parent

	shubham := User{"Shubham", "shubham@go.dev", true, 19}
	fmt.Println(shubham)
	fmt.Printf("shubham details are %+v\n", shubham)
	fmt.Printf("shubham Name is %v\n", shubham.Name)

	shubham.GetStatus()
	shubham.NewMail()
	fmt.Println(shubham)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("Email of this is user is:", u.Email)
}
