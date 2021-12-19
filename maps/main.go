package main

import "fmt"

func main() {
	/*colour := map[string]string{
		"red":  "colour1",
		"blue": "colour2",
	}*/

	//var colour map[string]string

	colour := make(map[int]string)

	colour[10] = "red"

	delete(colour, 10)

	colour = map[int]string{
		1: "color 1",
		2: "color 2",
	}
	printMap(colour)
}

func printMap(m map[int]string) {
	for k, v := range m {
		fmt.Println("key: ", k, ", Value: ", v)
	}
}
