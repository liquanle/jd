// timer project main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World!")

	//定时任务开始
	for {
		now := time.Now()
		fmt.Println("daka——", now.Format("2006-01-02 15:04:05"))
		// 计算下一个零点
		next := now.Add(time.Second * 70)
		next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), 0, 0, next.Location())
		fmt.Println("next daka——", next.Format("2006-01-02 15:04:05"))
		timer := time.NewTimer(next.Sub(now))
		<-timer.C
	}
	//定时任务结束
}
