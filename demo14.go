package main

import "fmt"

/**
 * @Author: whl
 * @Description:
 * @File:  demo14
 * @Date: 2023/04/05 20:28
 */
func main() {
	var a uint = 18
	var b uint = 66
	var c uint = 0
	//位运算
	c = a & b
	fmt.Printf("%d,二进制%b", c, c) //与
	fmt.Println()
	c = a | b
	fmt.Printf("%d,二进制%b", c, c) //或
	fmt.Println()
	c = a ^ b
	fmt.Printf("%d,二进制%b", c, c) //异或
	fmt.Println()
	c = a << 2
	fmt.Printf("%d,二进制%b", c, c) //整体左移
	fmt.Println()
	c = a >> 2
	fmt.Printf("%d,二进制%b", c, c) //整体右移
}
