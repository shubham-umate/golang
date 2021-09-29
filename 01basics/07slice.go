package main

import(
	"fmt"
)

func main(){
	var things = []string{"maleta","ropa","reloj"}
	fmt.Println(things)
	

	things = append(things, "bolso")
	fmt.Println(things)

	things = append(things[1:], )
	fmt.Println(things)
}