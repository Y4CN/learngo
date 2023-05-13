package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
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
	name := os.Args[1]
	l := utf8.RuneCountInString(name)
	name = strings.ToUpper(name)
	rep := strings.Repeat("!", l)
	s := rep + name + rep
	fmt.Println(s)
}
