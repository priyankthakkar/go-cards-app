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
