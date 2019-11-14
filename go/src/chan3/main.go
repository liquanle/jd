// chan3 project main.go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	fmt.Println("1")
	ch <- 3
	fmt.Println("oo", <-ch)
	fmt.Println(2)
	fmt.Println("Hello World!")
}
