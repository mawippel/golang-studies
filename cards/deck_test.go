package main

import "testing"

func NewDeckShouldHave16Cards(t *testing.T) {
	d := newDeck()

	deckLenght := len(d)
	if deckLenght != 16 {
		t.Errorf("Expected deck length of 16, but got %v", deckLenght)
	}
}

func FirstCardOfNewDeckShouldBeAceOfSpades(t *testing.T) {
	d := newDeck()

	firstDeckCard := d[0]
	if firstDeckCard != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", firstDeckCard)
	}

}

func LastCardOfNewDeckShouldBeFourOfClubs(t *testing.T) {
	d := newDeck()

	lastDeckCard := d[len(d)-1]
	if lastDeckCard != "Four of Clubs" {
		t.Errorf("Expected first card of Four of Clubs, but got %v", lastDeckCard)
	}
}
