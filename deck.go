package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func newDeck() (deck, error) {
	cards := deck{}
	cardSuites := []string{"Spade", "Diamond", "Heart", "Club"}
	cardValues := []string{"Ace", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "Jack", "King", "queen"}
	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	err := cards.toFile("deck")
	return cards, err
}

func (d deck) deal(handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) toFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0644)
}
