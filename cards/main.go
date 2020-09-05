package main

import "fmt"

func main() {
	cards := newDeck()
	firstPart, secondPart := deal(cards, 5)
	firstPart.print()
	secondPart.print()

	fmt.Println("toString: ", cards.toString())

	cards.saveToFile("my_cards")
}
