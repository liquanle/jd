// ticker project main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	myTicker := time.NewTicker(time.Second * 5) //定义时钟周期,每隔1秒写入时间

	i := 0 //记录次数

	ch := make(chan int)
	fmt.Println(k)

	go func() {
		fmt.Println("go1")
		ch <- 3
		fmt.Println("go2")
	}()

	time.Sleep(3 * time.Second)
	m := <-ch
	fmt.Println("m = ", m)
	fmt.Println("开始执行")
	for {
		<-myTicker.C //阻塞，形成间隔
		now := time.Now()
		curTime := now.Format("2006-01-02 15:04:05")
		i++
		fmt.Println(curTime)

		if i == 5 {
			myTicker.Stop()
			break
		}
	}

}
