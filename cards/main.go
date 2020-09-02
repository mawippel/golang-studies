package main

import "fmt"

func main() {
	cards := []string{"asd", "asd", "asd", "asd", "asd", "asd", newCard()}
	cards = append(cards, "aaaa")

	for _, card := range cards {
		fmt.Println(card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
