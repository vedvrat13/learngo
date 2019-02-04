package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Learn Interfaces
type bot interface {
	greeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	// eb := englishBot{}
	// sb := spanishBot{}

	// printGreeting(eb)
	// printGreeting(sb)

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, resp.Body)
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
