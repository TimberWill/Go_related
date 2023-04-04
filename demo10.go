package main

/**
 * @Author: whl
 * @Description: 数值类型
 * @File:  demo10
 * @Date: 2023/04/04 21:00
 */
import "fmt"

func main() {
	//定义一个整型
	var a int = 100
	fmt.Printf("%T,%d\n", a, a) //整数型对应的值，打印需要用到%d

	//定义一个浮点型
	var num float64 = 3.14
	fmt.Printf("%T,%f\n", num, num) //浮点型对应的值，打印需要用到%f
	fmt.Printf("%T,%.2f\n", num, num)
	fmt.Println(num)
}
