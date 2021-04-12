package main

import "fmt"

func show(c interface{}) {
	fmt.Println(c)
}

func showType(c interface{}) {
	v, ok := c.(string)
	if !ok {
		fmt.Println("类型转换错误")
	}
	fmt.Println(v)
}

func showType2(c interface{}) {
	switch c.(type) {
	case string:
		fmt.Printf("this c string type, c:%s \n", c)
		break
	case int:
		fmt.Printf("this c int type, c:%d \n", c)
		break
	case int64:
		fmt.Printf("this c int64 type, c:%d \n", c)
		break
	case bool:
		fmt.Printf("this c bool type, c:%v \n", c)
		break
	}
}

func main() {

	//map使用interface
	stuMap := make(map[string]interface{}, 10)
	stuMap["1"] = "sam"
	stuMap["2"] = 100
	fmt.Println(stuMap)

	//方法使用interface类型
	show("hello world")
	show(100)

	//类型断言
	showType(100)
	showType("1234as")
	showType2("123")
	showType2(false)
	showType2(int64(500))
}
