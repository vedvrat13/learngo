package main

func main() {
	cards := newDeckFromFile("new_file")

	// hand, remainingHands := deal(cards, 5)
	// hand.print()
	// fmt.Println("********")
	// remainingHands.print()
	// fmt.Println([]byte("Hi There"))
	// fmt.Println([]byte(cards.toString()))
	// cards.saveToFile("new_file")

	cards.shuffleDeck()
	cards.print()
}

func newCard() string {
	return "Five of Spades"
}
