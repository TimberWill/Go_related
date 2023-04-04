package main

/**
 * @Author: whl
 * @Description:匿名变量
 * @File:  demo5
 * @Version: 1.0.0
 * @Date: 2023/04/04 19:21
 */
import "fmt"

func test() (int, int) {
	return 100, 200
}

func main() {
	a, _ := test()
	fmt.Println(a)
}
