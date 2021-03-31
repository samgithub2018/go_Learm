package main

import "fmt"

//函数指针
//函数像变量一样传递

func main() {
	fmt.Println(f1(f2))
	fmt.Println(f3()())
	fmt.Println(f4())
}

//将函数作为形参
func f1(f func() int) int {
	return f()
}

func f2() int {
	return 5
}

//返回值作为参数
func f3() func() int {
	return f2
}

//测试返回值引用的关系
//defer(x) 是将x的引用就是内存地址传递进去，所以内存地址都一致
func f4() (x int) {
	defer func(int) {
		fmt.Println(&x)
		x++
		fmt.Println(&x)
	}(x)
	fmt.Println(&x)
	return 5
}


