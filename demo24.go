package main

import "fmt"

/**
 * @Author: whl
 * @Description: 打印九九乘法表
 * @File:  demo24
 * @Date: 2023/04/06 13:23
 */
func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
}
