package main

import (
	"fmt"
	"runtime"
)

func main()  {
	n := runtime.GOMAXPROCS(8) //指定一单核运行  设置运行的核数
	fmt.Println("n = ",n)

	for {
		go fmt.Print(1)

		fmt.Print(0)
	}

}
