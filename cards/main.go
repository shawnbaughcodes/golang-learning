package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("theCards")

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
