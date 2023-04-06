package main

import "fmt"

/**
 * @Author: whl
 * @Description: defer延迟执行
 * @File:  demo34
 * @Date: 2023/04/06 17:47
 */
func main() {
	a := 10
	fmt.Println("start a:", a)
	defer fn(a)
	a++
	fmt.Println("end a:", a)
}
func fn(s int) {
	fmt.Println("函数体中的a:", s)
}
