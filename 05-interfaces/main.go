package main

import "fmt"

type bot interface {
	// if there are a type  that had a function in this program called 'getGreeting' that return a string, then the type extends this interfaces. means that the type also can called all the function defined in this interface (in this case the 'printGreeting' function)
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
