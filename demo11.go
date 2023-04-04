package main

import "fmt"

func main() {
	var str string
	str = "hello,world"
	fmt.Printf("%T,%s\n", str, str)

	s1 := 'A'
	s2 := "A"
	s3 := '中'
	fmt.Printf("%T,%d\n", s1, s1) //整数型
	fmt.Printf("%T,%s\n", s2, s2)
	fmt.Printf("%T,%d\n", s3, s3) //整数型

	//转义符
	fmt.Println("hello," + "world")
	fmt.Println("hello,\"world")
	fmt.Println("hello,\nworld")
	fmt.Println("hello,\tworld")
}
