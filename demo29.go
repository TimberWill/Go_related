package main

import "fmt"

/**
 * @Author: whl
 * @Description: 可变参数
 * @File:  demo29
 * @Date: 2023/04/06 14:35
 */
func main() {
	getSum(1, 2, 3, 100)
}
func getSum(nums ...int) {
	sum := 0
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
		sum += nums[i]

	}
	fmt.Println("sum:", sum)
}
