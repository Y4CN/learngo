package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	//get CPU CORE number
	// fmt.Println(runtime.NumCPU())

	//StringConvTest
	// thisIsBool := "gg"
	// reallBool, e := strconv.ParseBool(thisIsBool)
	// fmt.Println(e)
	// fmt.Printf("this should return Boolian Type => %T , value is %t \n", reallBool, reallBool)

	//I love Practice
	// fArg := os.Args[1]
	// lArg := os.Args[2]
	// fFloatArg, _ := strconv.ParseFloat(fArg, 64)
	// lFloatArg, _ := strconv.ParseFloat(lArg, 64)
	// fmt.Println(fFloatArg + lFloatArg)

	//change C to F
	// cel := os.Args[1]
	// fCel, _ := strconv.ParseFloat(cel, 64)
	// far := fCel*1.8 + 32
	// fmt.Printf("%gF\n", far)

	// //get Lenghts String
	// name := "Y4CNش"
	// //this is show me len byte of the String :D
	// numberString := len(name)
	// //this Show me the reall Len of the String :D
	// reallnumber := utf8.RuneCountInString(name)
	// fmt.Println(numberString, reallnumber)

	//this is Repeat the !
	// name := os.Args[1]
	// l := utf8.RuneCountInString(name)
	// name = strings.ToUpper(name)
	// rep := strings.Repeat("!", l)
	// s := rep + name + rep
	// fmt.Println(s)

	//Create type
	// type Duration int64
	// var d Duration

	//if StateMent Chalenge
	// const userName1 = "Y4CN"
	// const userName2 = "Y4CN2"
	// const pass1 = "12345678"
	// const pass2 = "123456789"
	// fmt.Println(len(os.Args))
	// if len(os.Args) <= 2 {
	// 	fmt.Println("Plz Enter Your Username & Password")
	// 	return
	// }
	// if (os.Args[1] == userName1 && os.Args[2] == pass1) || (os.Args[1] == userName2 && os.Args[2] == pass2) {
	// 	if os.Args[1] == userName1 {
	// 		if os.Args[2] != pass1 {
	// 			fmt.Printf("Pass is invalid %s\n", userName1)
	// 			return
	// 		}
	// 		fmt.Printf("Welcome %s with a %s password\n", userName1, pass1)
	// 		return
	// 	}
	// 	if os.Args[1] == userName2 {
	// 		if os.Args[2] != pass2 {
	// 			fmt.Printf("Pass is invalid %s\n", userName2)
	// 			return
	// 		}
	// 		fmt.Printf("Welcome %s with a %s password\n", userName2, pass2)
	// 		return
	// 	}
	// } else {
	// 	fmt.Println("You Pass OR User is Invalid")
	// }
	// n, err := strconv.Atoi("22")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(n)
	// fmt.Printf("Type n is : %T\n", n)

	//chalenge Switch with time package
	// myTime := time.Now().Hour()
	// fmt.Println(myTime)
	// switch {
	// case myTime >= 0 && myTime < 10:
	// 	fmt.Println("Good Morning")
	// case myTime >= 10 && myTime < 15:
	// 	fmt.Println("Good Noon")
	// case myTime >= 15 && myTime < 21:
	// 	fmt.Println("Good Afternoon")
	// default:
	// 	fmt.Println("Good Night")
	// }

	//for Loop
	//Jadval Zard
	// const max = 15
	// fmt.Printf("%5s", "X")
	// for i := 0; i <= max; i++ {
	// 	fmt.Printf("%5d", i)
	// }

	// for i := 0; i <= max; i++ {
	// 	if i == 0 {
	// 		fmt.Println()
	// 	}
	// 	fmt.Printf("%5d", i)
	// 	for j := 0; j <= max; j++ {
	// 		fmt.Printf("%5d", i*j)
	// 	}
	// 	fmt.Println()
	// }

	//Lucky Number

	const turn = 5
	if len(os.Args) < 2 {
		fmt.Println("Plz Enter Number")
		return
	}
	guess, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Plz Enter a Number")
		return
	}
	if guess <= 0 {
		fmt.Println("Plz Enter Positive Number")
		return
	}
	//that was for debugging
	// guess := 6
	for i := 0; i < turn; i++ {
		n := rand.Intn(guess + 1)
		if n == guess {
			if i == 0 {
				fmt.Println("PASHMAAAAAAAAAAAAAAM .....")
				return
			}
			fmt.Println("Nice 😍")
			return
		}
	}
	fmt.Println("NOOOOO You Lost 😢")
}
