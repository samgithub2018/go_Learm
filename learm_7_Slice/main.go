package main

import "fmt"

func main() {

	//切片的定义
	SliceDefi()

	//切片的附加
	//s1 := []string{"广州", "上海", "北京"}
	//s1 = append(s1, "深圳", "佛山")
	//fmt.Println(s1)

	//测试切片的引用和扩充
	SliceCit()



}

//切片的定义
func SliceDefi() {
	//定义一个切片变量，不分配内存空间
	var s1 []int
	fmt.Println(s1)

	//初始化切片
	s1 = []int{1, 2, 3}
	fmt.Println(s1)

	a := [...]int{1, 2, 4, 6, 9, 7, 3, 10, 9, 14, 15}
	fmt.Printf("a: %d \n", len(a))

	//通过一个数组获得一个切片
	s2 := a[1:5]
	fmt.Println(s2)
	fmt.Printf("len(s2):%d cap(s2):%d \n", len(s2), cap(s2))
	//从数组的第一个元素切到第3个
	s3 := a[:3]
	fmt.Println(s3)
	fmt.Printf("len(s3):%d cap(s3):%d \n", len(s3), cap(s3))
	//从数组的第3个元素切到最后一个
	s4 := a[3:]
	fmt.Println(s4)
	fmt.Printf("len(s4):%d cap(s4):%d \n", len(s4), cap(s4))
	//切片再切片，根据上一个切片的容量初始化本切片的容量
	s5 := s4[0:5]
	fmt.Println(s5)
	fmt.Printf("len(s5):%d cap(s5):%d \n", len(s5), cap(s5))

	//测试切片的引用性，a数组的 10改成了 100
	s5[len(s5)-1] = 100
	fmt.Println(a)
	//通过修改底层数组 s5切片的10改成90
	a[7] = 90
	fmt.Println(s5)

	//使用make定义切片
	var m1 = make([]int, 5, 10)
	fmt.Println(m1)

	m1[3] = 100
	fmt.Println(m1)
}

//测试切片的引用和扩充
func SliceCit() {
	//定义一个固定容量的数组
	a := [15]int{1, 2, 3}
	fmt.Printf("len(a):%d cap(a):%d  \n", len(a), cap(a))

	//通过数组获得一个切片
	a1 := a[3:]
	fmt.Printf("len(a1):%d cap(a1):%d \n", len(a1), cap(a1))

	//给这个切片附加一个元素，并查看其容量变化
	a2 := append(a1, 9)
	fmt.Printf("len(a2):%d cap(a2):%d \n", len(a2), cap(a2))
	//测试是否还继续引用着a数组
	a2[0] = 1
	fmt.Println(a)
	fmt.Println(a1)
	fmt.Println(a2)

	//继续给a2切片添加容量，并观察a2 和 返回的a3的容量，发现其并扩充容量，还是继续使用a2底层的数组
	a3 := append(a2, 8)
	fmt.Printf("len(a2):%d cap(a2):%d \n", len(a2), cap(a2))
	fmt.Printf("len(a3):%d cap(a3):%d \n", len(a3), cap(a3))
	//通过查看这个就能得知a2和a3还是引用着
	a3[0] = 9
	fmt.Println(a2)
	fmt.Println(a3)
}
