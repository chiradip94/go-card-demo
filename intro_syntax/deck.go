package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
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

func readFile(f string) string {
	ar, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return string(ar)
}

func createDeckFromFile(fileName string) deck {
	deckStr := readFile(fileName)
	return deck(strings.Split(deckStr, ","))
}

func (d deck) shuffleDeck() deck {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := range d {
		randomNo := r1.Intn(len([]string(d)))
		d[i], d[randomNo] = d[randomNo], d[i]
	}
	return d
}
