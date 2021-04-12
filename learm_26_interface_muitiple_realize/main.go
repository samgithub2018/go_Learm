package main

import "fmt"

//基本动作
type basic interface {
	eat()
	move()
}

//人类
type person interface {
	talk()
	//继承了basic的接口，实例化person接口的时候，也必须实现basic里面的方法
	basic
}

//汉族
type hanzhu interface {
	speakChinese()
}

//程序员
type programmer struct {
}

//实现基本动作
func (p *programmer) eat() {
	fmt.Println("this programmer eat")
}
func (p *programmer) move() {
	fmt.Println("this programmer move")
}

//实现人类的动作
func (p *programmer) talk() {
	fmt.Println("this programmer talk")
}

//实现汉族的动作
func (p *programmer) speakChinese() {
	fmt.Println("你好golang")
}

func main() {

	//实例化基本动作接口
	var ba basic = &programmer{}
	ba.eat()
	ba.move()

	//实例化人类接口
	var per person = &programmer{}
	per.talk()
	ba = per //可以转成basic基类
	fmt.Printf("%T \n", ba) //*main.programmer

	//实例化汉族接口
	var han hanzhu = &programmer{}
	han.speakChinese()
	fmt.Printf("%T \n", han) //*main.programmer
}
