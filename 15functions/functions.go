package main

import "fmt"

func main() {
	fmt.Println("Lets learn functions")

	a := 10
	b := 20
	result := sum(a, b)
	fmt.Printf("Sum is %v", result)

	proRes, proStr := proAdder(2, 3, 4, 5)
	fmt.Printf("ProResilt is %v and String is %s", proRes, proStr)

}

func sum(a int, b int) int {
	return a + b
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi Pro result function"
}
