package main

import(
	"fmt"
)

func main(){

	var isLoggedIn bool = false
	var balance int = 10

	if isLoggedIn || balance > 5 {
		fmt.Println("Show cart page")
		var len, err = fmt.Println("Show cart page")
		if err == nil{
			fmt.Printf("length is %v, %T",len,len)
		}

	} else{
		fmt.Println("Show login page")
	}
}
	
