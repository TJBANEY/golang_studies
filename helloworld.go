package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// fmt.Println("======")
	// remainingCards.print()

	fmt.Println(cards.toString())
}
