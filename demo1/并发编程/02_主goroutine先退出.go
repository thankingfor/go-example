package main

import (
	"fmt"
	"time"
)

func main()  {

	go func() {
		i := 0
		fmt.Println("主协程 i = ", i)
		time.Sleep(time.Second)
	}()

	i := 0
	for {
		i++
		fmt.Println("main i = ", i)
		time.Sleep(time.Second)

		if i == 2 {
			break
		}
	}
}