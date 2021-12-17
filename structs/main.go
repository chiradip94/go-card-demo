package main

import "fmt"

type contactInfo struct {
	email   string
	pincode int
}

type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstname: "Jim",
		lastname:  "Pope",
		contact: contactInfo{
			email:   "abc@cde.com",
			pincode: 1234567,
		},
	}
	fmt.Printf("%+v\n", jim)
}
