package main

import "fmt"

// Create a new type of 'deck', which is a slice of string
type deck []string

// (d deck) is a receiver
// receiver = any variable of type 'deck' now gets access to the print method
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Hearts", "Spades", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, value := range cardValues {
		for _, suit := range cardSuits {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
