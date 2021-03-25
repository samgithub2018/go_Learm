package main

import "fmt"

const name = "sam"

const (
	a = 1
	b = 2
)

const (
	c = iota //索引从 0 开始 ，没加一个常量 iota + 1
	d = iota
	e = 100
	f = iota //3
)

func main() {

	fmt.Println(name)
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}
