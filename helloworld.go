package main

import "fmt"

func main() {
	cards := newDeckFromFile("my_cards")
	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// fmt.Println("======")
	// remainingCards.print()

	fmt.Println(cards.toString())
}
