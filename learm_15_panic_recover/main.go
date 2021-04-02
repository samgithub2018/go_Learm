package main

import "fmt"

func f2() {
	fmt.Println("程序执行开始")
	defer func() {
		//获取错误信息，并尝试修复让程序继续执行下去
		//也可以不接收，让程序就此终结
		err := recover()
		fmt.Println(err)
		fmt.Println("程序遇到错误，已修复完成")
	}()
	//弹出一个错误，可以是用recover接收
	panic("发生了严重的错误")
	fmt.Println("程序执行结束")
}

func main() {
	f2()
}
