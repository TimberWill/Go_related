package main

import "fmt"

/*
变量交换
*/
func main() {
	var a int = 100
	var b int = 200
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)

}
