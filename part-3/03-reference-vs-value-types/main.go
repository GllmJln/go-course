package main

import "fmt"

// Go is a pass-by value language, meaning when argument is passed to a function,
// the underlying value is passed, not the reference. For reasons explained in the course,
// Some types are actually "pass-by reference".

type structs struct {
	value string
}

func main() {
	//Reference types: these values should change
	slice := []string{"foo"}
	//maps discussed later
	maps := map[string]string{"foo": "bar"}
	//channels to be filled in
	pointers := &[]string{"bar"}
	//functions to be filled in
	updateReferenceTypes(slice, maps, pointers)
	fmt.Println(slice, maps, *pointers)

	//Value types: these values should not change
	ints := 12
	floats := 12.1
	strings := "bar"
	booleans := true
	structs := structs{
		value: "value",
	}
	valueTypes(ints, float32(floats), strings, booleans, structs)
	fmt.Println(ints, floats, strings, booleans, structs)
}

func updateReferenceTypes(slice []string, maps map[string]string, pointers *[]string) {
	slice[0] = "bar"
	maps["foo"] = "foobar"
	(*pointers)[0] = "foobar"
}

// This function does not do anything as these values are passed by value
func valueTypes(ints int, floats float32, strings string, booleans bool, structs structs) {
	ints = 78
	floats = 78.1
	strings = "foo"
	booleans = false
	structs.value = "foobar"
}
