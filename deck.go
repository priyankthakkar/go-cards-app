package main

import "fmt"

// this declares a new type called deck, and this new type is equivalent of
// a slice of type string
type deck []string

/**
* function newDeck() here, creates a new deck, adds multiple cards to and returns the deck
* we can see here, this function doesn't have receiver, because the aim of this function is
* to create a deck, not to operate on a already created deck
* we can see that this function has a return type, deck
 */
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Diamonds", "Spades", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

/**
* Let's talk about function print() now
* this method prints information about all the card within a deck
* it is define with keyworld func
* name of the function is print
* it doesn't have any parameters
* but what is (d deck) ?? Well, this portion withing the declration is called receiver
* What is represents is, function print() is a function on type deck
* it means any variable declared with type deck will have a function print() on it
* and when it calls the print() function, it is passed as an agrument to this function
* and is accessible via d, d is variable name and it can be absolutly any variable name
 */
func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}
