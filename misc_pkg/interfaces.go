package main

import "fmt"

//=========================//
//======= WRONG WAY =======//
//=========================//

// type englishBot struct{}
// type spanishBot struct{}

//=========================//
//
// Issue = getGreeting function most likely has very different logic inside of them, so two similar functions but with different recievers is appropriate
// The printGreeting function however is going to have identical logic regardless of if englishBot uses it or spanishBot uses it.
// How do you use the same function, but with two different custom types? => Interfaces
//
//=========================//

// If not making use of the receiver param at all, you can get rid of the alias
// func (englishBot) getGreeting() string {
// 	// Very custom logic for generating an English greeting
// 	return "Hi There"
// }

// func (spanishBot) getGreeting() string {
// 	// Very custom logic for generating a Spanish greeting
// 	return "Hola"
// }

// Error would be raised for uncommented out second printGreeting because you can't have perform function overloading in Go
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

//=========================//
//======= RIGHT WAY =======//
//=========================//

//========= INTERFACE REFACTOR =========//

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// Very custom logic for generating an English greeting
	return "Hi There"
}

func (spanishBot) getGreeting() string {
	// Very custom logic for generating a Spanish greeting
	return "Hola"
}

// STRUCTURE OF AN INTERFACE //
// 1 - "type bot interface" This broadcasts to the program "Hey! Everyone inside this program, all you different types out there, a new type is in town, the bot Interface type"
// 2 - "{ getGreeting() string }" If you are a type in this program with a function of getGreeting that returns a string, than you are now an honorary member of type 'bot'
// 3 - SpanishBot and EnglishBot are now also of type "Bot"
// 4 - Now that SpanishBot and EnglishBot are also of type Bot, they can be passed into any function with parameter Bot
// 4B - This is what allows one function with identical code to be used by two different types
// 3 - bot is just an arbitrary name
