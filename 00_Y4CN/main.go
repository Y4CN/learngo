package main

import (
	"fmt"
	"runtime"
)

func main() {
	//get CPU CORE number
	fmt.Println(runtime.NumCPU())

	//StringConvTest
	// thisIsBool := "gg"
	// reallBool, e := strconv.ParseBool(thisIsBool)
	// fmt.Println(e)
	// fmt.Printf("this should return Boolian Type => %T , value is %t \n", reallBool, reallBool)

}
