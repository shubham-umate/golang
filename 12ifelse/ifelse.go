package main

import "fmt"

func main() {
	fmt.Println("If else in golang")
	loginCount := 23
	var result string
	if loginCount < 10 {
		result = "Regular User"
	} else {
		result = "Something else"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is EVEN")
	} else {
		fmt.Println("Number is ODD")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is NOT less than 10")
	}
}
