package main

import "fmt"

/**
 * @Author: whl
 * @Description:匿名函数
 * @File: demo36
 * @Date: 2023/04/06 18:38
 */
func main() {
	//匿名函数
	f11 := func() {
		fmt.Println("我是函数f11")
	}
	//调用f11
	f11()

	//简化一下
	func() {
		fmt.Println("我是匿名函数")
	}()

	//可以传参数
	func(a, b int) {
		fmt.Println(a, b)
		fmt.Println("我是可以传参的匿名函数")
	}(1, 2)

	r1 := func(a, b int) int {
		fmt.Println("我是r1函数")
		return a + b
	}(1, 2)
	fmt.Println(r1)
}
