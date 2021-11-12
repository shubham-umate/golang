package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?courseName=reactjs&paymentid=lshfjh233jff"

func main() {

	fmt.Println("Welcome to urls in golang")

	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	
	qparams := result.Query()
	fmt.Printf("The type of query params are : %T\n",qparams)

	fmt.Println(qparams["courseName"])

	for _, v := range qparams {
		fmt.Println("Param is ",v)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "tutcss",
		RawPath: "user=shubham",
	}

	anotherURL := partsOfUrl.String()

	fmt.Println("Another URL ",anotherURL)

}
