package main

import "fmt"

/**
 * @Author: whl
 * @Description: 验证密码案例
 * @File:  demo18
 * @Date: 2023/04/05 21:15
 */
func main() {
	var first, second int
	var pwd int = 20230405

	//用户输入密码
	fmt.Println("请输入密码：")
	fmt.Scan(&first)

	//业务
	if first == pwd {
		fmt.Println("请再次输入密码:")
		fmt.Scan(&second)
		if second == pwd {
			fmt.Println("登录成功")
		} else {
			fmt.Println("登录失败，第二次密码错误")
		}
	} else {
		fmt.Println("登录失败")
	}
}
