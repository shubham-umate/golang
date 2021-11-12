package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"courseName"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags, omitempty"` // omitempty if tags value is null
}

func main() {
	fmt.Println("Welcome to JSON video")
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "lco.in", "abc123", []string{"web-dev", "js"}},
		{"Angular Bootcamp", 299, "lco.in", "abc123", []string{"web-dev", "js"}},
		{"Go Bootcamp", 299, "lco.in", "abc123", []string{"golang", "js"}},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFormWeb := []byte(`
	{
		"courseName": "Go Bootcamp",
		"price": 299,
		"website": "lco.in",
		"tags": ["golang","js"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFormWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFormWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON is not valid")
	}

	// in some cases you just want to add data in key value pairs (not using struct)

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFormWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is %T\n", k, v, v)
	}
}
