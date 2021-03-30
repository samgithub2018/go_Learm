package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

	//基本定义操作
	BasicDef()

	//
	rand.Seed(time.Now().UnixNano()) //创建随机数种子

	var m1 = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		va := rand.Intn(100)
		m1[key] = va
	}

	keys := make([]string, 200)
	for k, v := range m1 {
		keys = append(keys, k)
	}
	fmt.Println(keys)

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, m1[key])
	}

}

func BasicDef() {
	//定义一个map变量
	var m1 = make(map[string]int, 10)
	m1["北京"] = 10
	m1["广州"] = 20
	m1["上海"] = 21
	fmt.Println(m1)

	//获取map值
	v, ok := m1["广州"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("没有找到map值")
	}
	fmt.Println(m1["广州1"]) //不存在Key就会返回 value类型的默认值
	a := m1["广州1"]
	fmt.Println(a)

	//遍历map
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	//删除map元素
	delete(m1, "北京")
	delete(m1, "北京1") //即使是没有这个元素也不会报错
	fmt.Println(m1)
}
