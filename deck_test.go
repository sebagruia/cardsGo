package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected the length of deck to be 16, instead got %v", len(d))
	}
	if d[0] != "AceofSpades" {
		t.Errorf("First card is %v and expected AceofSpades", d[0])
	}
	if d[len(d)-1] != "FourofClubs" {
		t.Errorf("Last card is %v and expected FourofClubs", d[15])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, instead got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
