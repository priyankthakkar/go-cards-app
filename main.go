package main

import "fmt"

func main() {

	// now we can see, as the type deck is available within package main
	// we can access it and use it here
	// as type deck is equivalent of slice of string, we can treat type
	// deck like a slice of string and put new string elements to it
	cards := deck{"Ace of Diamond", newCard()}

	cards = append(cards, "Six of Spades")

	for index, card := range cards {
		fmt.Println(index, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
