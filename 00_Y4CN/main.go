package main

import (
	"fmt"
	"runtime"
)

func main() {
	//get CPU CORE number
	fmt.Println(runtime.NumCPU())
}
