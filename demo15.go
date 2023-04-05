package main

import "fmt"

/**
 * @Author: whl
 * @Description:
 * @File:  demo15
 * @Date: 2023/04/05 20:55
 */
func main() {
	//定义两个变量
	var x int
	var y float64

	fmt.Println("请输入两个数字：1.整数 2.浮点数")
	fmt.Scanln(&x, &y)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}
