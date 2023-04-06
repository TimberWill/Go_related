package main

import "fmt"

/**
 * @Author: whl
 * @Description: 引用传递
 * @File:  demo31
 * @Date: 2023/04/06 15:47
 */
func main() {
	/*
		引用传递
	*/
	//定义一个切片  []类型
	arr := []int{1, 2, 3, 4}
	fmt.Println("修改前arr", arr)
	//传递，拷贝arr
	update2(arr)
	fmt.Println("修改后arr:", arr)
}
func update2(arr1 []int) {
	arr1[0] = 100
	fmt.Println("arr1:", arr1)
}
