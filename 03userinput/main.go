package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to user input")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for Pizza:")
	// comma ok || Error ok syntax

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of rating is %T", input)

}
