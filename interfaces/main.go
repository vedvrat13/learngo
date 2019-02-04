package main

import "fmt"

type bot interface {
	greeting() string
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
	fmt.Println(b.greeting())
}

func (e englishBot) greeting() string {
	return "Hi There"
}

func (e spanishBot) greeting() string {
	return "Hola"
}
