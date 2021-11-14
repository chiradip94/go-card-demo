package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spade", "Diamond", "Heart", "Club"}
	cardValues := []string{"Ace", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "Jack", "King", "queen"}
	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) deal(handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}

}
