package main

import "fmt"

func main() {
	a := 100
	b := 100
	// 条件判断

	// if - else if - else
	if a > b {
		fmt.Println("a>b")
	} else if a == b {
		fmt.Println("a=b")
	} else {
		fmt.Println("a<b")
	}

	// switch 判断多个
	switch a {
	case 100:
		fmt.Println("100")
	case 200:
		fmt.Println("200")
	case 300:
		fmt.Println("300")
	default:
		// 如果上面的都不匹配, 就执行default
		fmt.Println("???")
	}

	// 循环
	// 定义i为0, 循环条件是i<5, 每次循环i+1
	for i := 0; i < 5; i++ {
		// 循环五次
		fmt.Println(i)
	}

	// 运行看看和上一个的不同之处
	for i := 0; i < 5; i++ {
		if i == 1 {
			fmt.Println("跳出当前这次循环, 但是不打断")
			continue
		}
		if i == 3 {
			fmt.Println("跳出循环, 打断循环")
			break
		}
		fmt.Println(i)
	}

	// 另一种
	i := 1

	for i < 5 {

		fmt.Println(i)

		i++
	}
}
