package main

import "fmt"

/**
 * @Author: whl
 * @Description: 打印5x5方阵
 * @File:  demo23
 * @Date: 2023/04/06 13:16
 */
func main() {
	for i := 0; i < 5; i++ {
		//打印行
		for j := 0; j < 5; j++ {
			fmt.Print("* ")
		}
		fmt.Println() //换行
	}
}
