package main

import "fmt"

// Create a new type if 'deck' which is a slice of strings

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
