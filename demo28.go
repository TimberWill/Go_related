package main

import "fmt"

/**
 * @Author: whl
 * @Description: 不同的函数类型
 * @File:  demo28
 * @Date: 2023/04/06 14:16
 */
func main() {
	printString()
	printA("hello,world!")
	add2(3, 8)
	k := add3(1, 2, 3)
	println(k)
	x, y := swap("hello", "whl")
	println(x, y)
}

// 无参无返回值函数
func printString() {
	fmt.Println("hello,world!")
}

// 有一个参数的函数
func printA(A string) {
	fmt.Println(A)
}

// 有两个参数的函数
func add2(a, b int) {
	c := a + b
	fmt.Println(c)
}

// 有一个返回值的函数
func add3(a, b, c int) int {
	d := a + b*c
	return d
}

// 有多个返回值的函数
func swap(x, y string) (string, string) {
	return y, x
}
