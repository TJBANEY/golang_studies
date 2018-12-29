package main

import (
	"fmt"

	codeWars "github.com/TJBANEY/golang_studies/code_wars_pkg"
)

func main() {
	neighbors := codeWars.CartesianNeighbor(2, 2)
	fmt.Println(neighbors)

	newNeighbors := codeWars.CartesianNeighborAlternate(2, 2)
	fmt.Println(newNeighbors)

	addressList := `John Daggett, 341 King Road, Plymouth MA
	                Alice Ford, 22 East Broadway, Richmond VA
					Sal Carpenter, 73 6th Street, Boston MA`

	organizedByState := codeWars.ByState(addressList)
	fmt.Println(organizedByState)

	// Create Players from player struct

	// tim := player{firstName: "Timothy", lastName: "Baney"}
	// fmt.Println(tim)

	// var alex player
	// alex.firstName = "Alex"
	// alex.lastName = "Test"
	// fmt.Printf("%+v", alex)

	// jim := player{
	// 	firstName: "Jim",
	// 	lastName:  "John",
	// 	contactInfo: contactInfo{
	// 		email:   "jim@gmail.com",
	// 		zipCode: 72758,
	// 	},
	// }

	// This won't work because Golang is a 'Pass by Value' language meaning that the receiver in newFirstName is going to copy jim value at RAM address 1, and save the copy at another RAM address
	// jim.newFirstName("Jimmy")
	// jim.printSelf()

	// We need to create a pointer to jim's address in memory, and update the receiver in newFirstName to accept a pointer as a parameter using an asterisk before the value
	// jimPointer := &jim
	// jimPointer.newFirstNamePointer("Jimmy")
	// jim.printSelf()

	// print RAM address value for curiousity and attempt printing address fetch for pointer, and non-pointer values.

	// Should return RAM address
	// fmt.Printf("%+v", jimPointer)

	// Should return value of jimPointer RAM address value
	// fmt.Printf("%+v", *jimPointer)

	// Not sure what this will return, since jim is a struct, not a RAM address - Tested and it does raise an error
	// fmt.Printf("%+v", *jim)

	// creating new pointer variable to pass in to function receiver is a little tedious, and there is a shortcut for it.
	// As long as the type is correct, you can pass in a non pointer value to a function receiver expecting a pointer value.
	// Golang will automatically convert "player type" into "player pointer type" for us.
	// jim.newFirstNamePointer("Blurbington")
	// jim.printSelf()

	// ======== MAPS SANDBOX ========= //

	// There is more than one way to declare a Go map

	// 1 - Plain Simple Literal Syntax
	// map[string]string <- Brackets state that all keys are type string, string outside of brackets state all values are of type string
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ffff",
	// }

	// 2 - Similar syntax but using var, and not assigning value
	// Typically use this when you want to figure out later on what key-value pairs we want to add to it.
	// var colorsTwo map[string]string

	// 3 - Using Go built-in function 'make'
	// colorsThree := make(map[string]string)

	// fmt.Printf("%+v", colors)
	// fmt.Printf("%+v", colorsTwo)
	// fmt.Printf("%+v", colorsThree)

	// To access key you must use bracket notation. You can't use dot notation.
	// All keys inside of map are 'Typed' so you can't use dot notation since the keys can be things other than strings.
	// fmt.Println(colors["red"])
	// colors["blue"] = "#0000ff"
	// delete(colors, "red")

	// Iterate over map
	// printMap(colors)

	// When to use Struct, and when to use Map since they are similar.
	// 1 - Keys and values in Map are strongly typed, but they are not in Structs
	// 2 - Keys are NOT indexed in Structs
	// 3 - Map is reference type, and Struct is value type
	// 4 - With Maps, you don't need to know all the keys at compile time. You can create it, delete fields, add fields, change fields before code is compiled.
	// 5 - With Structs, you NEED to know all the keys at compile time.
	// Use Map when you need to represent a collection of related properties
	// Use Struct when you need to represent a "Thing" with a lot of different properties, much like a JS object literal, or Python class

	// ====== INTERFACES SANDBOX ======= //

	// Since function receivers can only be one type, if we wanted to use the same function but with a different type, we would need to create a brand new function
	// Interfaces solves this problem

	// eb := englishBot{}
	// sb := spanishBot{}

	// printGreeting(eb)
	// printGreeting(sb)

	// makeGetRequest()
	// c := make(chan string)
	// iterateUrls(c)
}

// receiver will copy to new RAM address here
// func (p player) newFirstName(newName string) {
// 	p.firstName = newName
// }

// func (p *player) newFirstNamePointer(newName string) {
// 	(*p).firstName = newName
// }

// func (p player) printSelf() {
// 	fmt.Printf("%+v", p)
// }

func cardStuff() {
	//======= Card tutorial stuff to be deleted/sorted later ========//

	// cards := newDeckFromFile("my_cards")
	// cards.saveToFile("my_cards")

	// cards := newDeck()
	// cards.shuffle()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// fmt.Println("======")
	// remainingCards.print()
	// fmt.Println(cards.toString())
}
