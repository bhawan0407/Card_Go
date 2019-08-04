package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected length of 52, but got %d", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected first card as Ace of Clubs, but got %s", d[0])
	}

	if d[51] != "King of Hearts" {
		t.Errorf("Expected last card as King of Hearts, but got %s", d[51])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {

	// delete the testing files (if any) to provide a clean testing environment
	os.Remove("_decktesting")

	cards := newDeck()
	cards.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected length of 52, but got %d", len(loadedDeck))
	}

	// delete the testing files after the test ran
	os.Remove("_decktesting")
}
