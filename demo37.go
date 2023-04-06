package main

import "fmt"

/**
 * @Author: whl
 * @Description: 回调函数
 * @File: demo37
 * @Date: 2023/04/06 19:05
 */
func main() {
	re1 := operate(1, 2, ad) //传入ad函数
	fmt.Println(re1)

	re2 := operate(3, 4, sub) //传入sub函数
	fmt.Println(re2)

	//传入匿名函数
	re3 := operate(9, 3, func(i int, j int) int {
		if j != 0 {
			return i / j
		} else {
			fmt.Println("除数不能为0")
			return 0
		}
	})
	fmt.Println(re3)
}

// 高阶函数,可以接收一个函数作为参数
func operate(a, b int, fun func(int, int) int) int {
	result := fun(a, b)
	return result
}

// 加法
func ad(a, b int) int {
	return a + b
}

// 减法
func sub(a, b int) int {
	return a - b
}
