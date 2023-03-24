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

	jimPointer := &jim // & operator points to the address in memory
	jimPointer.updateName("jimmy")
	// The above can be shortened to 
	jim.updateName("jimmy")
	// This is a shortcut, however notice the type mismatch between jim and jimPointer -> this is allowed

	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) { //* on a type means is somewhat different to * on a var
	(*pointerToPerson).firstName = newFirstName //* gives access to the value siting at the address in memory
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
