package main

import "fmt"

func main() {
	card, err := newDeck()
	if err != nil {
		fmt.Println(err)
	}
	hand, card := card.deal(4)
	fmt.Println(hand)
}
