package main

func main() {

	// Much, much cleaner main() function, the call to function newDeck() here
	// creates a new deck of cards and it is being assigned to cards
	cards := newDeck()

	// we know that function print() is declared with receiver (d deck)
	// it means cards which is of type deck, has a function print() in them
	// and when cards invoke the function print(), instance of cards will be passed
	// to function print() via variable d in the receiver declration
	// which will print information about all the cards
	cards.print()
}
