package main

func main() {
	cards := deck{"asd", "asd", "asd", "asd", "asd", "asd", newCard()}
	cards = append(cards, "aaaa")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
