package main

import "fmt"

/**
 * @Author: whl
 * @Description: 函数
 * @File:  demo27
 * @Date: 2023/04/06 14:02
 */
func main() {
	fmt.Println("hello,world!")
	fmt.Println(add(1, 2))
}

func add(a, b int) int {
	c := a + b
	return c
}
