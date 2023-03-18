package main

import "fmt"

func main() {
	cards := []string{"Ace of Spades", "Nine of Clubs"}
	cards = append(cards, "Six of Diamonds") //Append does not modify the cards var

	print(cards)

}

func print(c []string) { //Don't actually know if thats the proper func arg syntax yet
	for i, card := range c {
		fmt.Println(i, card)
	}
}
