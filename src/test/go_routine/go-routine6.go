package main

import (
	"fmt"
	"time"
)

var ch chan int

func test_channel() {

	ch <- 1
	fmt.Println("ch 1")
	ch <- 1
	fmt.Println("ch 2")
	ch <- 1
	fmt.Println("time to end goroutine 1")
}

func main() {

	//	ch = make(chan int)  //等价于ch = make(chan int,0)
	ch = make(chan int, 2) //不会阻塞
	go test_channel()
	time.Sleep(2 * time.Second)
	fmt.Println("running end!")
	<-ch
	time.Sleep(time.Second)
}
