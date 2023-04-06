package main

import "fmt"

/**
 * @Author: whl
 * @Description: fallthrough一路穿透到底
 * @File:  demo21
 * @Date: 2023/04/06 11:01
 */
func main() {
	var score int = 90
	switch score {
	case 90:
		fmt.Println("A")
		fallthrough
	case 80:
		if score == 90 {
			break
		}
		fmt.Println("B")
		fallthrough
	case 60, 70:
		fmt.Println("C")
		fallthrough
	default:
		fmt.Println("D")
	}
}
