package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

func (cards deck) toString() string {
	return strings.Join([]string(cards), ",")
}

func deal(d deck, numberOfCards int) (deck, deck) {
	return d[:numberOfCards], d[numberOfCards:]
}

func (cards deck) saveToFile(filename string) error {
	cardsByteArray := []byte(cards.toString())
	return ioutil.WriteFile(filename, cardsByteArray, 0666)
}
