package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of Deck
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{
		"Spades", "Diamonds", "Harts", "Clubs",
	}
	cardValues := []string{
		"Ace", "Two", "Three", "Four",
	}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+"of"+cardSuit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toByte() []byte {
	deckstring := strings.Join(d, ",")
	fmt.Println(deckstring)
	return []byte(deckstring)
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, d.toByte(), 0644)
}
