package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func newDeckFromFile(filename string) deck {
	bs, error := ioutil.ReadFile(filename)
	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(0)
	}
	stringSlice := strings.Split(string(bs), ",")
	return deck(stringSlice)
}

func (cards deck) shuffle() {
	randomSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range cards {
		newPosition := randomSource.Intn(len(cards) - 1)
		cards[i], cards[newPosition] = cards[newPosition], cards[i]
	}
}
