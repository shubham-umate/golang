package main

import "fmt"

func main() {
	fmt.Println("Maps by golang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all langguages ", languages)
	fmt.Println("JS ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	// loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
