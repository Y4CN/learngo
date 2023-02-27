// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"math/rand"
	"os"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Moodly
//
//   1. Get username from command-line
//
//   2. Display the usage if the username is missing
//
//   3. Create an array
//     1. Add three positive mood messages
//     2. Add three negative mood messages
//
//   4. Randomly select and print one of the mood messages
//
// EXPECTED OUTPUT
//
//   go run main.go
//     [your name]
//
//   go run main.go Socrates
//     Socrates feels good 👍
//
//   go run main.go Socrates
//     Socrates feels bad 👎
//
//   go run main.go Socrates
//     Socrates feels sad 😞
//
//   go run main.go Socrates
//     Socrates feels happy 😀
//
//   go run main.go Socrates
//     Socrates feels awesome 😎
//
//   go run main.go Socrates
//     Socrates feels terrible 😩
// ---------------------------------------------------------

func main() {
	rand.Seed(time.Now().UnixNano())

	moods := [...]string{
		"good 👍",
		"bad 👎",
		"sad 😞",
		"happy 😀",
		"awesome 😎",
		"terrible 😩",
	}

	if len(os.Args) != 2 {
		println("[your name]")
		return
	}

	name := os.Args[1]
	mood := moods[rand.Intn(6)]
	println(name, "feels", mood)
}
