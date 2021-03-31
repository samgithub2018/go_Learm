package main

import "fmt"

func f1(x int) {
	fmt.Printf("x=%d \n", x)
}

func f2() int {
	fmt.Println("f2")
	return 5
}

// 执行顺序是
// start
// defer f1(f2()) f1的参数会先拿到，所以会先执行f2()拿到形参
// end
// x=5
func f3() {
	fmt.Println("start")
	defer f1(f2())
	fmt.Println("end")
}

func main() {
	fmt.Println("f3")
	f3()
	fmt.Println("f2")
	f5()
	fmt.Println("f6")
	f6()
	fmt.Println("f7")
	f7()
}

func f4(s string, a, b int, ) {
	fmt.Println(s, a, b)
}

//说说打印的顺序
// c 0 1
// s 0 2
// a 1 2
//得出的结论是defer func如果有参数的时候会先拿到【当时的】具体值，而不是引用
func f5() {
	a := 1
	b := 2
	defer f4("a", a, b)
	a = 0
	defer f4("b", a, b)
	b = 1
	defer f4("c", a, b)
}


//说说打印的顺序
// c 0 1
// b 0 2
// a 1 2
//得出的结论是defer和前面的f5 一样
func f6() {
	a := 1
	b := 2
	defer func(q, w int) {
		fmt.Println("a", q, w)
	}(a, b)
	a = 0
	defer func(q, w int) {
		fmt.Println("b", q, w)
	}(a, b)
	b = 1
	defer func(q, w int) {
		fmt.Println("c", q, w)
	}(a, b)
}


//说说打印的顺序
// c 0 1
// b 0 1
// a 0 1
//得出的结论是，如果匿名函数中的匿名函数使用“外层函数”的变量，会是引用关系
//像是成员变量一样会一直存储到所有的匿名函数执行完成
func f7() {
	a := 1
	b := 2
	defer func() {
		fmt.Println("a", a, b)
	}()
	a = 0
	defer func() {
		fmt.Println("a", a, b)
	}()
	b = 1
	defer func() {
		fmt.Println("a", a, b)
	}()
}
