package main

import "fmt"

/**
 * @Author: whl
 * @Description: defer延迟
 * @File:  demo33
 * @Date: 2023/04/06 16:19
 */
func main() {
	a(1)
	a(2)
	defer a(3)
	a(4)
	a(5)
	defer a(6)
	a(7)
	defer a(8)
	a(9)
	defer a(10)
}
func a(num int) {
	fmt.Println(num)
}
