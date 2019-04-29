package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	joel := person{
		firstName: "Joel",
		lastName:  "McDaniel",
		contactInfo: contactInfo{
			email:   "joelmcdbme@gmail.com",
			zipCode: 83835,
		},
	}

	joel.updateName("Geoli")
	joel.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
