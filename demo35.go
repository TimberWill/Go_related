package main

import "fmt"

/**
 * @Author: whl
 * @Description:函数的数据类型
 * @File: demo35
 * @Date: 2023/04/06 18:28
 */
func main() {
	fmt.Printf("%T\n", f1)
	fmt.Printf("%T\n", f2)
	fmt.Printf("%T\n", "hello")
	fmt.Printf("%T\n", 100)
	//定义函数类型变量
	var f4 func(int, int)
	f4 = f3
	fmt.Println(f4)
	fmt.Println(f3) //f4和f3地址值相同
	f4(1, 2)

}
func f1() {

}
func f2(a, b int) int {
	return 1
}
func f3(a, b int) {
	fmt.Println(a, b)
}
