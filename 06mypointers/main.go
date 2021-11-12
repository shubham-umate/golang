package main

import "fmt"

func main() {
	fmt.Println("Let's learn pointers")

	num := 10

	var ptr = &num

	fmt.Println("Value of actual ptr is: ", ptr)
	fmt.Println("Value of *ptr is: ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("Value of actual ptr is: ", ptr)
	fmt.Println("Value of *ptr is: ", *ptr)

}
