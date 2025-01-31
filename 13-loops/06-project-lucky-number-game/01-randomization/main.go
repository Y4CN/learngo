// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// rand.Seed(10)
	// rand.Seed(100)

	// t := time.Now()
	// rand.Seed(t.UnixNano())

	// ^-- same:

	//we can Use random type without Seed now this is Depricated

	// rand.Seed(time.Now().UnixNano())

	guess := 10

	for n := 0; n != guess; {
		// fmt.Println("here is Rand Number ", rand.Intn(10))
		n = rand.Intn(guess + 1)
		fmt.Printf("%d\n ", n)
	}
	fmt.Println()
}
