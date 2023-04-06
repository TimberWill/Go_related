package main

import "fmt"

/**
 * @Author: whl
 * @Description:值传递
 * @File:  demo30
 * @Date: 2023/04/06 15:22
 */
func main() {
	/*
		值传递
	*/
	//定义一个数组  [个数]类型
	arr := [4]int{1, 2, 3, 4}
	fmt.Println("修改前arr", arr)
	//传递，拷贝arr
	update(arr)
	fmt.Println("修改后arr:", arr)

}
func update(arr1 [4]int) {
	arr1[0] = 100
	fmt.Println("arr1:", arr1)
}
