package main

import "fmt"

/**
 * @Author: whl
 * @Description: switch语句练习
 * @File:  demo19
 * @Date: 2023/04/06 10:43
 */
func main() {
	var score int = 98

	//匹配
	switch score {
	case 90:
		fmt.Println("A")
	case 80:
		fmt.Println("B")
	case 60, 70:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}

	switch { //默认的条件：bool = true
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("true")
	default:
		fmt.Println("其他")
	}
}
