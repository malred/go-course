package main

import "fmt"

func main() {
	// 条件判断
	a := 100
	b := 100
	// if elseif else
	if a > b {
		fmt.Println("a>b")
	} else if a == b {
		fmt.Println("a=b")
	} else {
		fmt.Println("a<b")
	}

	// switch 多个条件
	switch a {
	case 100:
		fmt.Println("a=100")
	case 200:
		fmt.Println("a=200")
	default:
		// 如果其他的都不匹配, 就走default
		fmt.Println("???")
	}

	// 循环

	// 定义i, 如果i<5就继续循环, 每次循环后i+1
	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}

	// 另一种写法
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// break & continue
	for i := 0; i < 5; i++ {
		if i == 1 {
			fmt.Println("跳出当前这一次循环, 但是不打断")
			continue
		}
		if i == 3 {
			fmt.Println("跳出循环, 打断循环")
			break
		}
		fmt.Println(i)
	}
}
