package main

import "fmt"

var value1 string

func main() {
	card := newDeck()
	hand := deck{}
	fmt.Print(hand)
	hand, card = card.deal(4)
	hand.print()
}
