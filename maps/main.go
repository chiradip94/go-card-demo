package main

import "fmt"

func main() {
	/*colour := map[string]string{
		"red":  "colour1",
		"blue": "colour2",
	}*/

	//var colour map[string]string

	colour := make(map[10]string)

	colour[10] = "red"

	fmt.Println(colour)
}
