package main

import (
	"fmt"
)

/**
 * @Author: whl
 * @Description: string
 * @File:  demo26
 * @Date: 2023/04/06 13:39
 */
func main() {
	str := "hello,world!"

	//获取字符串长度
	fmt.Println("打印字符串长度：", len(str))
	//获取某一位对应的字符
	fmt.Println("打印字符串第2位对应的字符：", str[1]) //打印该位字符对应的ascii码
	fmt.Printf("%c\n", str[1])
	fmt.Println("====================================")
	//遍历打印string字符串
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}
	fmt.Println("------------------------------------")

	//特殊的for循环，针对切片方法
	//打印的是字符串中的下标，对应的字符
	for i, v := range str {
		fmt.Print(i)
		fmt.Printf("%c\n", v)
	}
}
