package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// INTERFACE FUNCTION
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//  CUSTOM BOT FUNCTIONALITIES
func (englishBot) getGreeting() string {
	return "Hi there"
}
func (spanishBot) getGreeting() string {
	return "Hola"
}
