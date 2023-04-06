package main

import "fmt"

/**
 * @Author: whl
 * @Description: break and continue
 * @File:  demo25
 * @Date: 2023/04/06 13:31
 */
func main() {
	//break
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("=============================================")
	//continue
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
