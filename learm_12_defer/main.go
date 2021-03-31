package main

import "fmt"

func f1() {
	fmt.Println("start")
	defer fmt.Println("aaaaa") //延迟执行到 函数执行完成
	fmt.Println("end")
}

//返回 5，因为x++是在返回值赋值之后才执行的
func f2() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

//返回6，因为x就是返回值变量，最终不管怎么执行都是返回x的值
func f3() (x int) {
	x = 5
	defer func() {
		x++
	}()
	return x
}

//返回5，因为x是值类型，y是在x++之前执行的
func f4() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

//就近原则，defer func(x) 中的x是形参，不具备引用关系，所以不管怎么改也不会影响到f5(x)的返回值x
func f5() (x int) {
	x=5
	defer func(x int) {
		x++
	}(x)
	return  x
}

func main() {
	f1()
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println(f5())
}
