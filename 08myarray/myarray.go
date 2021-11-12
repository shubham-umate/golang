package main

import "fmt"

func main() {
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[2] = "Tomato"

	fmt.Println(fruitList)

	fmt.Println(len(fruitList))

	var veggieList = [5]string{"Potato", "beans"}
	fmt.Println(len(veggieList))
}
