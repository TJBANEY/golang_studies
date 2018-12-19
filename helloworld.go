package main

import (
	"fmt"
)

func main() {
	// cards := newDeckFromFile("my_cards")
	// cards.saveToFile("my_cards")

	cards := newDeck()
	cards.shuffle()

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

	// print RAM address value for curiousity and attempt printing address fetch for pointer, and non-pointer values.

	// Should return RAM address
	fmt.Printf("%+v", jimPointer)

	// Should return value of jimPointer RAM address value
	fmt.Printf("%+v", *jimPointer)

	// Not sure what this will return, since jim is a struct, not a RAM address - Tested and it does raise an error
	// fmt.Printf("%+v", *jim)

	// creating new pointer variable to pass in to function receiver is a little tedious, and there is a shortcut for it.
	// As long as the type is correct, you can pass in a non pointer value to a function receiver expecting a pointer value.
	// Golang will automatically convert "player type" into "player pointer type" for us.
	jim.newFirstNamePointer("Blurbington")
	jim.printSelf()

	// ======== MAPS SANDBOX ========= //

	// There is more than one way to declare a Go map

	// 1 - Plain Simple Literal Syntax
	// map[string]string <- Brackets state that all keys are type string, string outside of brackets state all values are of type string
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ffff",
	}

	// 2 - Similar syntax but using var, and not assigning value
	// Typically use this when you want to figure out later on what key-value pairs we want to add to it.
	var colorsTwo map[string]string

	// 3 - Using Go built-in function 'make'
	colorsThree := make(map[string]string)

	fmt.Printf("%+v", colors)
	fmt.Printf("%+v", colorsTwo)
	fmt.Printf("%+v", colorsThree)

	// To access key you must use bracket notation. You can't use dot notation.
	// All keys inside of map are 'Typed' so you can't use dot notation since the keys can be things other than strings.
	fmt.Println(colors["red"])
	colors["blue"] = "#0000ff"
	delete(colors, "red")

	// When to use Struct, and when to use Map since they are similar.
}

// receiver will copy to new RAM address here
func (p player) newFirstName(newName string) {
	p.firstName = newName
}

func (p *player) newFirstNamePointer(newName string) {
	(*p).firstName = newName
}

func (p player) printSelf() {
	fmt.Printf("%+v", p)
}
