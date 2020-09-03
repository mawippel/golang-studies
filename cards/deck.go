package main

import "fmt"

// create a new type of deck which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (cards deck) print() {
	for _, card := range cards {
		fmt.Println(card)
	}
}

func deal(d deck, numberOfCards int) (deck, deck) {
	return d[:numberOfCards], d[numberOfCards:]
}
