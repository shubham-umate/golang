package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video")
	//PerformGetRequest()
	//PerformPostRequest()
	PerformPostFormRequest()

}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Response Status Code: ", response.StatusCode)

	fmt.Println("Response Content Length: ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	// This is response string
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte count is", byteCount)
	fmt.Println(responseString.String())

	//fmt.Println("Content :", string(content))
}

func PerformPostRequest() {

	const myurl = "http://localhost:8000/post"

	// fake json payload

	requestBody := strings.NewReader(`
		{
			"courseName":"Let's go with the golang",
			"price":0,
			"platform":"lco"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("Response is", string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"
	// form data

	data := url.Values{}

	data.Add("firstName", "shubham")
	data.Add("lastName", "umate")
	data.Add("email", "sh@go.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
	defer response.Body.Close()
}
