package main

import "fmt"

func main() {
	cards := []string{}
	cards = append(cards, "Six of Spades")

	for _, card := range cards {
		fmt.Println(card)
	}
}

func newCard() string {
	return "Ace of Spades"
}
