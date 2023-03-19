package main

import "fmt"

type contactInfo struct {
	email    string
	postCode int
}

type niceStuff struct {
	noticeTheKeyInPersonStruct int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	niceStuff
}

func main() {
	// Alex and John arw two ways of defining structs
	alex := person{
		"Alex",
		"Anderson",
		contactInfo{
			"alex@gmail.com",
			389204,
		},
		niceStuff{1},
	}

	john := person{
		firstName: "John",
		lastName:  "Smith",
		contact: contactInfo{
			email:    "john@email.com",
			postCode: 3679248,
		},
	}

	fmt.Printf("%+v\n%+v", alex, john)

	var sarah person
	fmt.Println(sarah)
	fmt.Printf("%+v", sarah)
	sarah.firstName, sarah.lastName = "Sarah", "Smith"

	//Receiver functions on structs
	sarah.print()
	sarah.updateName("Mel")
	sarah.print()
}

func (p person) updateName(newFirstName string) { // This does not update the person's name, doing that part next
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
