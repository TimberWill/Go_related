package main

import "fmt"

/**
 * @Author: whl
 * @Description:fallthrough穿透
 * @File:  demo20
 * @Date: 2023/04/06 10:56
 */
func main() {
	a := false
	switch a {
	case false:
		fmt.Println("1. case条件为false")
		fallthrough
	case true:
		fmt.Println("2. case条件为true")
	}
}
