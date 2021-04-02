package main

import "fmt"

//闭包

//这里的a已经引用起来，被第一次calc调用的时候缓存起来了，通过看打印的三个&a
//对于f1和f2来说a 就是一个成员变量，所以f1和f2每次调用都会改变其值，且是累计的
func calc(a int) (func(int) int, func(int) int) {
	fmt.Println(&a)
	add := func(i int) int {
		fmt.Println(&a)
		a += i
		return a
	}
	sub := func(i int) int {
		fmt.Println(&a)
		a -= i
		return a
	}
	return add, sub
}

func main() {
	f1, f2 := calc(10)

	a1 := f1(1)
	a2 := f2(2)

	//返回值不是引用的，而是在内存中复制一份
	fmt.Printf("a1=%x \n",&a1)
	fmt.Printf("a2=%x \n",&a2)

	fmt.Println(a1, a2) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7
}
