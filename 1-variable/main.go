package main

import "fmt"

// 程序入口
func main() {
	// 不指定类型
	var v1 = 100
	var v2 = "hello"

	// 不赋初始值
	var v3 int
	var v4 string
	var v5 bool

	// 赋值
	v4 = "hhh"

	// 多个相同类型赋值
	var v6, v7 int = 1, 2

	// 多个不同类型赋值
	var v8, v9 = 1, "??"

	// 短变量声明+初始化 (我用的比较多)
	v10 := 100
	v11 := "abc"

	// 同时 相同类型
	v12, v13 := 200, 300

	// 同时 不同类型
	v14, v15 := 114514, "abc"

	// 控制台输出
	fmt.Println(v1, v2)
	fmt.Println(v3, v4, v5)
	fmt.Println(v6, v7)
	fmt.Println(v8, v9)
	fmt.Println(v10, v11)
	fmt.Println(v12, v13)
	fmt.Println(v12, v13)
	fmt.Println(v14, v15)
	// 变量定义了不使用会报错
}
