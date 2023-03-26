package main

import "fmt"

type bot interface {
	getGreeting() string
}

// Notice how interfaces are implicit, there is no "implements" or anything like that
type spanishBot struct{}
type englishBot struct{}

func main() {
	eb := englishBot{}
	sp := spanishBot{}
	printGreeting(eb)
	printGreeting(sp)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}
func (spanishBot) getGreeting() string {
	return "Hola!"
}
