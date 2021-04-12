package main

import "fmt"

//【1】接口类型 定义了一个call方法
type animal interface {
	call()
	eat(string2 string) //传一个参数的函数
}

type dog struct{}

type cat struct{}

//【2】dog实现了 animal的call方法
func (d dog) call() {
	fmt.Println("dog")
}
func (d dog) eat(s string) {
	fmt.Printf("doc eat %s \n", s)
}

//【2】cat实现了 animal的call方法
func (c cat) call() {
	fmt.Println("call")
}
func (c cat) eat(s string) {
	fmt.Printf("cat eat %s \n", s)
}

//【3】 调用call方法，传实现了接口的方法的结构体（结构体的特定方法实现
func call(a animal) {
	a.call()
	a.eat("meal")
}

func main() {
	call(dog{})
	call(cat{})

	//接口的动态类型化
	var an animal
	fmt.Printf("%T \n", an) //<nil>

	an = dog{}
	fmt.Printf("%T \n", an) //main.dog
	an.call()

	an = cat{}
	fmt.Printf("%T \n", an) //main.cat
	an.call()
}
