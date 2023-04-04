package main

import "fmt"

func main() {
	var isFlag bool = true
	var isFlag2 bool = false
	fmt.Println(isFlag, isFlag2)

	fmt.Printf("%T,%t\n", isFlag, isFlag)
	fmt.Printf("%T,%t\n", isFlag2, isFlag2)
}
