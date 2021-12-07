package main

import "testing"

func TestNewDeck(t *testing.T) {
	deck1, _ := newDeck()

	if len(deck1) != 52 {
		t.Errorf("Expected 52, got %v", len(deck1))
	}
}
