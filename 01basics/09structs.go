package main

import (
	"fmt"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	rob := User{"rob", "rob@lco.dev", 23}
	fmt.Printf("%+v\n", rob)
	fmt.Printf("%+v\n", rob.Email)

	var sam = new(User)
	sam.Name = "Sam"
	sam.Email = "sam@lco.dev"
	sam.Age = 22

	fmt.Printf("%+v\n", sam)

	var tobby = &User{"tobby", "tobby@lco.dev", 22}
	fmt.Printf("%+v", tobby)

}
