package main

import "fmt"

/**
 * @Author: whl
 * @Description: 闭包结构
 * @File: demo38
 * @Date: 2023/04/06 19:28
 */
func main() {
	r1 := increment()
	fmt.Println(r1()) //这里打印的是这个函数的地址
	//要先调用，才会返回对应的值,给定一个变量来接收
	v1 := r1()
	fmt.Println(v1)
	v2 := r1()
	fmt.Println(v2)
	v3 := r1()
	fmt.Println(v3)
	//新定义一个r2函数指向increment函数
	r2 := increment()
	v4 := r2()
	fmt.Println(v4)
	v5 := r1()
	fmt.Println(v5) //再次调用r1，局部变量没有被释放
	v6 := r1()
	fmt.Println(v6)
}

// 自增,返回函数
func increment() func() int {
	//局部变量i
	i := 0
	//定义一个匿名函数，给变量自增并返回
	fun := func() int {
		i++
		return i
	}
	return fun
}
