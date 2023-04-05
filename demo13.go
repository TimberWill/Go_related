package main

import "fmt"

/**
 * @Author: whl
 * @Description:
 * @File:  demo13
 * @Date: 2023/04/05 20:11
 */
func main() {
	var a int = 10
	var b int = 3

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a % b)
	a++
	fmt.Println(a)
	b--
	fmt.Println(b)
}
