package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	e := 52
	if len(d) != e {
		t.Errorf("Expected deck length of %v, but got %v", e, len(d))
	}

	ex := "Ace of Spades"
	if d[0] != ex {
		t.Errorf("Expected first card of %v, but got %v", ex, d[0])
	}

	ex = "King of Clubs"
	if d[len(d)-1] != ex {
		t.Errorf("Expected first card of %v, but got %v", ex, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	f := "_decktesting"
	os.Remove(f)

	d := newDeck()
	d.saveToFile(f)

	loadedDeck := newDeckFromFile(f)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards but got %v", len(loadedDeck))
	}

	os.Remove(f)
}
