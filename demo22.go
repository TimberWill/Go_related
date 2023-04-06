package main

import "fmt"

/**
 * @Author: whl
 * @Description: for循环
 * @File:  demo22
 * @Date: 2023/04/06 11:10
 */
func main() {
	//循环输出1-10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//输出1-10的数字之和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i

	}
	fmt.Println(sum)

	//将循环初始条件写到外面，自增自减写到里面
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}
