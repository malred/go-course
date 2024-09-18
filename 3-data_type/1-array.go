package main

import "fmt"

func main() {
	// 定义数组
	var arr1 [3]string
	//arr2 := [3]string{"a", "b", "c"}
	//arr2 := []string{"a", "b", "c"} // 也可以不指定长度
	arr2 := [...]string{"a", "b", "c"}
	// 数组的索引从0开始 0 1 2 3 4
	arr3 := [5]int{1: 10, 3: 30} // 1和3的位置赋值, 其他位置的值为默认

	fmt.Println(arr1, arr2, arr3)
	fmt.Println()

	// 选取数组中的某段
	fmt.Println(arr3[:])   // 全部
	fmt.Println(arr3[1:3]) // 从1到2, 不包含3
	fmt.Println(arr3[:3])  // 从0到2, 不包含3
	fmt.Println()

	// 循环遍历数组
	i := 0
	for range arr3 {
		fmt.Println(arr3[i])
		i++
	}
	fmt.Println()
	// i为当前索引, v接收arr中当前的值
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	fmt.Println()
	// 不需要的值可以用_代替
	for _, v := range arr3 {
		fmt.Println(v)
	}
	fmt.Println()

	// 复制值
	arr5 := arr3  // 复制了值, 改变arr5, 不会改变arr3
	arr6 := &arr3 // 复制了地址, 改变arr6, 会改变arr3的值

	arr5[1] = 50
	fmt.Println(arr5, arr3)
	arr6[1] = 60
	fmt.Println(arr6, arr3)

	// arr6打印出来是&, 也就是指针的形式, 如何获取指针对应的值? -> *
	fmt.Println(*arr6)
}
