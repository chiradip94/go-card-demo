package main

import "fmt"

type englishBot struct {
	name string
}
type banglaBot struct {
	name string
}

type commonGreet interface {
	greet() string
}

func main() {

	e := englishBot{}
	b := banglaBot{}
	callMe(e)
	callMe(b)
}

func (englishBot) greet() string {
	return "Hello"
}

func (banglaBot) greet() string {
	return "Nomoskar"
}

func callMe(g commonGreet) {
	fmt.Println(g.greet())
}
