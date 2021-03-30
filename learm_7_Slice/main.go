package main

import "fmt"

func main() {

	//切片的定义
	SliceDefi()

	//测试切片的引用和扩充
	SliceCit()

	//删除元素
	DelElement()


	
}

//切片的定义
func SliceDefi() {
	//定义一个切片变量，不分配内存空间
	var s1 []int
	fmt.Println(s1)

	//初始化切片
	s1 = []int{1, 2, 3}
	fmt.Println(s1)

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

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

//测试切片的引用和扩容
func SliceCit() {

	//测试容量小于1024，扩容元素小于2倍以内的元素，是否会翻倍扩容容量
	a1 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("len(a1)=%d, cap(a1)=%d \n", len(a1), cap(a1)) //len(a1)=7, cap(a1)=7
	a2 := append(a1, 1, 2, 3)
	fmt.Printf("len(a2)=%d, cap(a2)=%d \n", len(a2), cap(a2)) //len(a2)=10, cap(a2)=14

	//测试容量小于1024，且扩容添加的元素大于原数组容量的两倍是否会根据需要的容量作为实际容量
	s1 := []int{1, 2, 3}
	fmt.Printf("len(s1)=%d, cap(s1)=%d \n", len(s1), cap(s1)) //len(s1)=3, cap(s1)=3
	s2 := append(s1, 4, 5, 6, 7, 8, 9)
	fmt.Printf("len(s2)=%d, cap(s2)=%d \n", len(s2), cap(s2)) //len(s2)=9, cap(s2)=10 ，容量当遇到奇数+1

	//测试容量大于1024的扩容机制，是否会根据文档所述以25%增加
	d1 := [1025]int{}
	for i := 0; i < 1025; i++ {
		d1[i] = i
	}
	fmt.Printf("len(d1)=%d, cap(d1)=%d \n", len(d1), cap(d1)) //len(d1)=1025, cap(d1)=1025
	d2 := d1[1:]
	fmt.Printf("len(d2)=%d, cap(d2)=%d \n", len(d2), cap(d2)) //len(d2)=1024, cap(d2)=1024
	d3 := append(d2, 1026)
	fmt.Printf("len(d3)=%d, cap(d3)=%d \n", len(d3), cap(d3)) //len(d3)=1025, cap(d3)=1280 差不多是这个数

	//测试切片引用的机制，在达不到新建数组情况下，还会继续引用原数组
	//通过这个实验得出结论，如果底层数组的容量够用则不会重新创建数组，而是继续使用原数组
	f1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	f2 := append(f1, 10)
	f3 := append(f2, 11)
	f3[0] = 10
	fmt.Printf("f1=%v  len(d3)=%d, cap(d3)=%d \n", f1, len(f1), cap(f1)) // f1=[1 2 3 4 5 6 7 8 9]  len(d3)=9, cap(d3)=9
	fmt.Printf("f2=%v  len(d3)=%d, cap(d3)=%d \n", f2, len(f2), cap(f2)) // f2=[10 2 3 4 5 6 7 8 9 10]  len(d3)=10, cap(d3)=18
	fmt.Printf("f3=%v  len(d3)=%d, cap(d3)=%d \n", f3, len(f3), cap(f3)) // f3=[10 2 3 4 5 6 7 8 9 10 11]  len(d3)=11, cap(d3)=18

	//两个切片合并
	g1 := []int{12, 13, 15}
	g2 := []int{16, 18, 19}
	g1 = append(g1, g2...)
	fmt.Println(g1)
}

//删除元素
func  DelElement()  {
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//把上面a数组中的索引2，3的值删了
	s1 := append(a1[:2], a1[3:]...)
	fmt.Println(s1) //[1 2 4 5 6 7 8 9]
	fmt.Println(a1) //[1 2 4 5 6 7 8 9 9]

	//其实现的原理是
	a2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := a2[:2]
	for i := 3; i < len(a2); i++ {
		s2 = append(s2, a2[i])
	}
	s2[0] = 100//测试是否还是继续引用原数组
	fmt.Println(s2)
	fmt.Println(a2)

	//得出结论是，修改通过修改切片，那么就相当于修改原数组
	//注意：数组的容量是不可变的，所以最多只是将数据的位置移动一下，或者将元素的值改掉

	//查看内存地址
	fmt.Println(&s2[0])
	fmt.Println(&a2[0])
	fmt.Println(&s2[3])
	fmt.Println(&a2[3])

}