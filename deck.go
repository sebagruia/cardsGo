package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r:= rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d)-1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}

func newDeckFromFile(fileName string) deck {
	cardsByteSlice, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	ss := strings.Split(string(cardsByteSlice), ",")
	return deck(ss)
}
