package main

import "fmt"

func main() {
	name := "whl"
	age := 24

	fmt.Println(name, age)
	//查看变量类型,注意是printf函数
	fmt.Printf("%T,%T", name, age)

}
