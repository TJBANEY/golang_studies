package main

import "fmt"

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
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 72758,
		},
	}

	fmt.Printf("%+v", jim)
}
