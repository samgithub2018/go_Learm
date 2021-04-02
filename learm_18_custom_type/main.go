package main

import "fmt"

//自定义类型
type myInt int

//类型别名
type youInt = int

func main() {
	var m myInt
	fmt.Printf("%T \n", m)
	m = 100
	fmt.Printf("%v \n", m)

	var y youInt
	fmt.Printf("%T \n", y)
	y = 100
	fmt.Printf("%v \n", y)
}
