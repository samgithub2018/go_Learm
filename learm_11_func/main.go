package main

import "fmt"

func main() {
	f1()
	f2(9, 10)

	//参数为一个切片，切片类型必须放在参数的最后
	f3("hello", 10, 11, 12, 13, 14)

	//带返回值的函数
	a := f4()
	fmt.Printf("a:%d\n", a)

	//多个返回值的函数
	b, c := f5()
	fmt.Println(b, c)

	//带命名的返回参数
	e, f := f6()
	fmt.Println(e, f)

}

func f1() {
	fmt.Println("f1 func")
}

func f2(x, y int) {
	fmt.Printf("x+y=%d\n", x+y)
}

//参数为切片
func f3(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}

//带返回值的函数
func f4() int {
	return 9 + 10
}

//多个返回值的函数
func f5() (int, string) {
	return 1, "1"
}

//带命名的返回参数
func f6() (ref1, ref2 int) {
	ref1 = 1
	ref2 = 2
	return
}
