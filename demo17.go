package main

import "fmt"

/**
 * @Author: whl
 * @Description: 分数评级
 * @File:  demo17
 * @Date: 2023/04/05 21:09
 */
func main() {
	var score int = 92

	if score >= 90 && score <= 100 {
		fmt.Println("A")
	} else if score >= 80 && score <= 90 {
		fmt.Println("B")
	} else if score >= 70 && score <= 80 {
		fmt.Println("C")
	} else if score >= 60 && score <= 70 {
		fmt.Println("D")
	} else {
		fmt.Println("不及格")
	}
}
