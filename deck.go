package main

import "fmt"

// Create a new type of 'deck', which is a slice of string
type deck []string

// (d deck) is a receiver
func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}
