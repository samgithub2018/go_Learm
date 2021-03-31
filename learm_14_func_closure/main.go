package main

import "fmt"

//闭包

//这里的 a已经引用，被第一次calc调用的时候缓存起来了
//对于f1和f2来说a 就是一个成员变量，所以f1和f2每次调用都会改变其值，且是累计的
func calc(a int) (func(int) int, func(int) int) {
	add := func(i int) int {
		a += i
		return a
	}
	sub := func(i int) int {
		a -= i
		return a
	}
	return add, sub
}

func main() {
	f1, f2 := calc(10)

	fmt.Println(f1(1), f2(2)) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7
}
