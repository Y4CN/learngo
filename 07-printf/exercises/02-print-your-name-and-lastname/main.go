// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print Your Name and LastName
//
//  Print your name and lastname using Printf
//
// EXPECTED OUTPUT
//  My name is Inanc and my lastname is Gumus.
//
// BONUS
//  Store the formatting specifier (first argument of Printf)
//    in a variable.
//  Then pass it to printf
// ---------------------------------------------------------

func main() {
	// BONUS: Use a variable for the format specifier

	// f := "My name is %v and my lastname is %v.\n"
	// fmt.Printf(f, "Yaode", "Kuang")
	fName := "Yasin"
	lName := "varaste"
	fmt.Printf("My First Name is %s & my Last Name is %s\n", fName, lName)
}
