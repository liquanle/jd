package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var printChar chan int

func prinNums() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		fmt.Println("prinNums", i)
		printChar <- 1111
		fmt.Println("printnum", <-printChar)
	}
}

func printChars() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		fmt.Println("printchar阻1", i)
		fmt.Println("printChars", <-printChar)
		fmt.Println("printchar阻2", i)
		fmt.Println("出来1")
		printChar <- 1222
		fmt.Println("出来2")
	}
}

func main() {
	printChar = make(chan int)

	wg.Add(2)

	go prinNums()
	go printChars()

	wg.Wait()
}
