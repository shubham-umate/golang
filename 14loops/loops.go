package main

import "fmt"

func main() {
	fmt.Println("Welcome to the loops in golang")

	//days := []string{"Sunday", "Monday", "Tuesday", "Wednesday","Friday", "Thursday"}

	// for i := 0; i < len(days); i++ {
	// 	fmt.Printf("Value at index %v is %v\n", i,days[i])
	// }

	// for _, day := range days {
	// 	fmt.Printf("Value at index  is %v\n", day)
	// }

	// for i, day := range days {
	// 	fmt.Printf("Value at index %v is %v\n", i, day)
	// }

	// for i := range days {
	// 	fmt.Printf("Value at index %v is %v\n", i, days[i])
	// }

	temp := 1

	for temp < 10 {

		// if temp == 5 {
		// 	//break
		// 	temp++ // While using continue we need to increment of decrement the value on which loop is running otherwise that iteration will get skipped and value won't be changed,
		// 	// so it will be an infinite loop
		// 	continue
		// }

		if temp == 5 {
			goto testGoto
		}

		fmt.Printf("Value is %v\n", temp)
		temp++
	}

testGoto:
	fmt.Println(" Jumping out of loop with goto")

}
