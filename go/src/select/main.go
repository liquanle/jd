package main

import (
	"fmt"
	"time"
)

func doneSelect(ch chan int) {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("for循环-输出-start")
		select {
		case data := <-ch:
			fmt.Println(data)
			break // 使用break只会跳出select
			//return //使用return会跳出for循环
		default:
			fmt.Println("select default testing")
		}
		fmt.Println("for循环-结束-end")
	}
	fmt.Println("for循环-跳出-end")
}

func do() {
	ch := make(chan int)
	go doneSelect(ch)
	ch <- 666
}

func main() {
	do()
	time.Sleep(3 * time.Second)
}
