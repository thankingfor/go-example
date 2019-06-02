package main

import "fmt"

func main()  {
	array := []int{0,1,2,3,4,5,6,7,8,9}
	//[low:hight:max] 取下标从low开始的元素 len=high - low cap = max - low
	s1 := array[:]//[0:len(array):len(array)]
	prin(s1)

	//操作莫格元素和数组操作方式一样
	data := array[1]
	fmt.Println("data = ",data)

	s2 := array[3:6:7]  //a[3] a[4] a[5]  len=6-3=3  cap=7-3=4
	prin(s2)

	s3 := array[:6] //从0开始  cap = len = 6
	prin(s3)

	s4 := array[3:] //从3开始到结尾
	prin(s4)
}

func prin(s []int)  {
	fmt.Println("s = ",s)
	fmt.Printf("len = %d ,cap = %d\n",len(s),cap(s))
}