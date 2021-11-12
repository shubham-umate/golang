package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Switch and case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Printf("Dice value is %v\n", diceNumber)
	case 2:
		fmt.Printf("Dice value is %v\n", diceNumber)
	case 3:
		fmt.Printf("Dice value is %v\n", diceNumber)
		fallthrough
	case 4:
		fmt.Printf("Dice value is %v\n", diceNumber)
		fallthrough
	case 5:
		fmt.Printf("Dice value is %v\n", diceNumber)
	case 6:
		fmt.Printf("Dice value is %v and roll the dice again\n", diceNumber)
	}
}
