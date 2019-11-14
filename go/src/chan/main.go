// chan project main.go
package main

import (
	"fmt"
	"time"
	//"time"
)

func main() {
	ch := make(chan int)
	// go func() {
	// 	for i := 0; i < 5; i++ {
	// 		time.Sleep(time.Duration(5-i) * time.Second)
	// 		fmt.Println("go inner 前", i)
	// 		ch <- i
	// 		fmt.Println("go inner 后", i)
	// 	}
	// }()

	go test1(ch)

	// for j := 0; j < 5; j++ {
	// 	fmt.Println("chan 前")
	// 	oa := <-ch
	// 	fmt.Println("chan 后", oa)
	// 	fmt.Println()
	// }
	time.Sleep(2 * time.Second)
	//a := <-ch
	//fmt.Println(a)
	time.Sleep(6 * time.Second)

	fmt.Println("end")
}

func test1(c chan int) {
	fmt.Println("test1 前")
	c <- 35 //放进去以后直接阻塞
	fmt.Println("test1 后")
}
