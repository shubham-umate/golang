package main

import(
	"fmt"
)

func main(){
	// make - allocates memory as well as initialzes memory
	// new - only allocates memory

	score := make(map[string]int) // uses new internally
	score["shubham"] = 89
	score["kiran"] = 84
	score["suraj"] = 98
	score["tushar"] = 89
	fmt.Println(score)

	getTusharScore := score["tushar"]
	fmt.Println(getTusharScore)

	delete(score,"kiran")
	fmt.Println(score)


	for k, v := range score{
		fmt.Printf("Score of %v is %v\n",k,v)
	}
}