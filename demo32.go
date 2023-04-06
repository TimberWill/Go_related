package main

import "fmt"

/**
 * @Author: whl
 * @Description: 递归函数
 * @File:  demo32
 * @Date: 2023/04/06 16:13
 */
func main() {
	sum := getSumN(5)
	fmt.Println(sum)
}

// getSumN(5) = getSumN(4)+5 = getSumN(3)+4+5 = getSumN(2)+3+4+5 = getSumN(1)+2+3+4+5 = 1+2+3+4+5
func getSumN(n int) int {
	if n == 1 {
		return 1
	}
	return getSumN(n-1) + n
}
