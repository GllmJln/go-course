package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Smith",
	}

	jim.print()
	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
