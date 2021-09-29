package main

import (
	"fmt"
)

func main(){
	var numbers [3]string
	numbers[0] = "uno"
	numbers[1] = "dos"
	numbers[2] = "tres"

	fmt.Println(numbers)

	var colors = [4]string{"roho","gris","azul","verde"}
	fmt.Println(colors)
	fmt.Println(colors[3])
	fmt.Println(len(colors))
}