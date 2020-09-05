package main

import "fmt"

func main() {
	cards := newDeck()
	firstPart, secondPart := deal(cards, 5)
	firstPart.print()
	secondPart.print()
	fmt.Println("toString: ", cards.toString())

	printSeparator()

	cards.saveToFile("my_cards")
	cardsFromFile := newDeckFromFile("my_cards")
	cardsFromFile.print()

	printSeparator()

	cards.shuffle()
	cards.print()
}

func printSeparator() {
	fmt.Println("==========================")
}
