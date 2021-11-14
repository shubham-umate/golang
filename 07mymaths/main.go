package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	//"math/rand"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	// var numOne int = 1
	// var numTwo float64 = 2.5

	// fmt.Println(numOne + int(numTwo))

	// random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	// random from crypto

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))

	fmt.Println(myRandomNum)
}
