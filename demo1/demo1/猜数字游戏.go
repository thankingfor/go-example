package main

import (
	"fmt"
	"math/rand"
)
import "time"

func main(){
	var randNum  int
	CreatNum(&randNum)

	fmt.Print("randNum： ",randNum)
	randSlice := make([]int ,4)
	//保存这个4位数的每一位
	GetNum(randSlice,randNum)
	fmt.Println("randSlice = ",randSlice)
	OnGame(randSlice)
}


func OnGame(randSlice []int)  {
	var num int
	for {
		fmt.Println("请输入一个4位数：")
		fmt.Scan(&num)
		if num <10000 && num >999 {
			break
		}
		fmt.Println("不符合要求：")
	}
	fmt.Println("num  =  ",num)

	keySlice := make([]int,4)
	GetNum(keySlice,num)
	fmt.Println("keySlient = ",keySlice)
}

func GetNum(s []int, num int) {
	//取千位
	s[0] = num / 1000
	s[1] = num % 1000 / 100
	s[2] = num % 100 / 10
	s[3] = num % 10
}

func CreatNum(p *int) {
	//设置种子
	rand.Seed(time.Now().UnixNano())

	var num int
	for {
		num = rand.Intn(10000)
		if num >= 1000 {
			break
		}
	}
	//fmt.Println("num = ",num)
	*p = num
}
