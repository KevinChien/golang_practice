package main

import "fmt"

func add(a int, b int) {
	c := a + b
	fmt.Println(c)
}

// if you need return something
func add_2(a int, b int) int {
	c := a + b
	return c
}

func add_3(a int, b int) int {
	return a + b
}

func main() {

	/*
		f := add_3() ==>應該是add_3不是add_3()
		var result int
		result = f(10, 20)
		fmt.Println(result)
	*/

	// 将函数赋值给变量 f
	f := add_3
	// 使用变量 f 调用函数并传递参数
	result := f(10, 20)
	// 打印结果
	fmt.Println(result)
}
