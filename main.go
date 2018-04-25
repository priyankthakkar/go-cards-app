package main

import "fmt"

func main() {
	// 1. First or an explicit way to declared a variable
	// It is important to remember that GO is strictly typed language
	// var card string = "Ace of spades"

	// 2. Shorter way of declaring variables, we can ommit the world
	// var as well as the datatype, because we know at complite time itself
	// what type of value is being assinged.
	card := "Ace of spades"
	fmt.Println(card)

	// one more important thing to notice is, := symbol is used only at the
	// time of variable declaration, after that only equla = symbol is used
	card = "Five of Hearts"
	fmt.Println(card)

	// we are calling function newCard() here which returns string
	card = newCard()
	fmt.Println(card)

	printState()

	/**
	* GO has two basic data structures, array and slice
	* array and slice are strongly typed, so while defining them datatypes are must
	* at a time, they can store only one type of items
	* array is fixed in length, so while defining them size is must
	* slice are similar to array but they are dynamic in size, they can shrink and grow
	 */

	// defining a slice of string with name cards
	// we can add string directly or we can all a function which returns string
	cards := []string{"Ace of Diamonds", newCard()}
	fmt.Println(cards)

	// to add new elements to slice we can use append function
	// append doesn't modify the exisiting slice
	// it takes an old slice as an input, takes the element to be added and
	// returns a new slice
	cards = append(cards, "Six of Spades")
	fmt.Println(cards)

	/**
	* to iterate over a slice of cards here, we are using for loop
	* here range represents the slice, while index represents index of a card in the slice
	* and aCard is a single card during the iteration
	* here we are printing both index & aCard, it is important to know that
	* every variable declared within for loop or inside program must be used
	 */
	for index, aCard := range cards {
		fmt.Println(index, aCard)
	}
}

/**
* newCard() here is function which returns a string, the structer is
* 1st func keyword to declare the function
* 2nd function name which is newCard
* and 3rd string which depcits the return type of the function
 */
func newCard() string {
	return "Five of Diamonds"
}
