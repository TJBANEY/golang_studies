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

	// Create Players from player struct
	tim := player{firstName: "Timothy", lastName: "Baney"}
	fmt.Println(tim)

	// fmt.Println(cards.toString())
}
