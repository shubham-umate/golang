package main

import (
	"fmt"
)

func main(){
	var score int = 34
	p := &score

	if p != nil{
		fmt.Println("P is having a value: ", *p)
	} else {
		fmt.Println("Watchout for nil values")
	}
}