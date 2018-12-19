package main

import (
	"fmt"
)

func main() {
	// cards := newDeckFromFile("my_cards")
	cards := newDeck()
	cards.shuffle()
	// cards.saveToFile("my_cards")

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// fmt.Println("======")
	// remainingCards.print()

	// fmt.Println(cards.toString())

	// Create Players from player struct

	// tim := player{firstName: "Timothy", lastName: "Baney"}
	// fmt.Println(tim)

	var alex player
	alex.firstName = "Alex"
	alex.lastName = "Test"
	fmt.Printf("%+v", alex)

	jim := player{
		firstName: "Jim",
		lastName:  "John",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 72758,
		},
	}

	// This won't work because Golang is a 'Pass by Value' language meaning that the receiver in newFirstName is going to copy jim value at RAM address 1, and save the copy at another RAM address
	jim.newFirstName("Jimmy")
	jim.printSelf()

	// We need to create a pointer to jim's address in memory, and update the receiver in newFirstName to accept a pointer as a parameter using an asterisk before the value
	jimPointer := &jim
	jimPointer.newFirstNamePointer("Jimmy")
	jim.printSelf()
}

// receiver will copy to new RAM address here
func (p player) newFirstName(newName string) {
	p.firstName = newName
}

func (p *player) newFirstNamePointer(newName string) {
	p.firstName = newName
}

func (p player) printSelf() {
	fmt.Printf("%+v", p)
}
