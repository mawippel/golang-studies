package main

func main() {
	cards := newDeck()
	//cards.print()
	firstPart, secondPart := deal(cards, 5)
	firstPart.print()
	secondPart.print()
}
