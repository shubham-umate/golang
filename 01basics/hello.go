package main

import "fmt"

func main() {
	var batman string = "I am batman"
	fmt.Println(batman)

	var superman string
	superman = "I am superman"

	fmt.Println(superman)

	thor := "I am thor"

	fmt.Println(thor)

	//thorRating := 3
	//fmt.Printf("%v,  %T",thorRating,thorRating)

	var ironMan, capAmerica string = "I am Ironman", "I am CapAmerica"

	fmt.Println(ironMan, capAmerica)

	var (
		spiderman = "I am spiderman"
		age       = 18
		powers    = "web slings, spidy sense"
		antman    = "I am antman"
	)

	fmt.Println(spiderman,age,powers,antman)
}
