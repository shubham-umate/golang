package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Goroutine")

	go greeter("Hello")
	greeter("World")
}

func greeter(s string) {

	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println(s)
	}

}
