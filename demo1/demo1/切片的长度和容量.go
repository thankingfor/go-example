package main

import "fmt"

func main(){
	a := []int{1,2,3,0,0}

	s := a[0:3:5]

	fmt.Println("S = ",s)
	fmt.Println("len(s) = ",len(s)) //切片的长度
	fmt.Println("cap(s) = ",cap(s)) //切片的容量

	y := a[1:4:5]

	fmt.Println("S = ",y)
	fmt.Println("len(s) = ",len(y)) //切片的长度
	fmt.Println("cap(s) = ",cap(y)) //切片的容量
}
