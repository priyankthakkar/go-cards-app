package main

func main() {

	// now we can see, as the type deck is available within package main
	// we can access it and use it here
	// as type deck is equivalent of slice of string, we can treat type
	// deck like a slice of string and put new string elements to it
	cards := deck{"Ace of Diamond", newCard()}

	cards = append(cards, "Six of Spades")

	// we know that function print() is declared with receiver (d deck)
	// it means cards which is of type deck, has a function print() in them
	// and when cards invoke the function print(), instance of cards will be passed
	// to function print() via variable d in the receiver declration
	// which will print information about all the cards
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
