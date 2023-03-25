package main

import (
	"fmt"
)

func main() {
	// creating maps
	colours := map[string]string{
		"red":   "#ff0000",
		"blue":  "#00ff00",
		"green": "#0000ff",
	}
	var colors map[string]string
	coloors := make(map[string]string)
	fmt.Println(colours, colors, coloors)

	// setting values
	colours["white"] = "#ffffff"
	fmt.Println(colours)

	// getting values - can't use the dot syntax because the key is typed
	white := colours["white"]
	fmt.Println(white)

	// deleting keys
	delete(colours, "white")
	fmt.Println(colours)

	printMap(colours)
}

// can iterate over the key, value pairs, unlike with structs
func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println(color, hex)
	}
}
