package main

import (
	"fmt"
)

func main(){
	rating := 3
	switch rating {
	case 1:
		fmt.Println("One")
	case 2,3:
		fmt.Println("Two or three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Other")
	}
}