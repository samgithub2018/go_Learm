package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func main() {

	//基本定义操作
	BasicDef()

	// 排序字符串
	MapSort()

	//切片中的map
	//map中的切片
	MapSlice()

	//获取一段字符串中字母，在字符串中的个数
	StrControls()



}

//基本定义操作
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

// 排序字符串
func MapSort() {

	//创建map字符串
	rand.Seed(time.Now().UnixNano()) //创建随机数种子
	var m1 = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		va := rand.Intn(100)
		key := fmt.Sprintf("stu%d", va)
		m1[key] = i
	}
	//拿出所有map的key
	keys := make([]string, 0, 200)
	for k, _ := range m1 {
		keys = append(keys, k)
	}
	fmt.Println(keys)
	//将map的key排序
	sort.Strings(keys)
	//通过for遍历排序
	for _, key := range keys {
		fmt.Println(key, m1[key])
	}
}

//切片中的map
//map中的切片
func MapSlice() {
	//切片中的map
	//定义一个切片
	s1 := make([]map[int]string, 0, 10)
	//定义一个map
	m1 := make(map[int]string, 3)
	m1[0] = "上海"
	m1[1] = "上海"
	m1[2] = "上海"
	//切片中添加map
	s1 = append(s1, m1)
	fmt.Println(s1)

	//map中的切片
	m2 := make(map[int][]string, 10)
	s2 := make([]string, 3)
	s2[0] = "上海"
	s2[1] = "上海"
	s2[2] = "上海"
	m2[0] = s2
	s3 := make([]string, 3)
	s3[0] = "北京"
	s3[1] = "北京"
	s3[2] = "北京"
	m2[1] = s3
	fmt.Println(m2)

}

//获取一段字符串中字母，在字符串中的个数
func StrControls() {
	s := "hello do you"
	strs := strings.Split(s, "")
	fmt.Println(strs)

	m1 := make(map[string]int, len(strs))
	for _, k := range strs {
		if k == " " {
			continue
		}
		_, ok := m1[k]
		if ok {
			m1[k] += 1
		} else {
			m1[k] = 1
		}
	}
	fmt.Println(m1)
}
