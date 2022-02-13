package main

import (
	"fmt"
	"time"
)

func send(ch chan int, e int) {
	for {
		ch <- e
		time.Sleep(time.Second)
	}
}

func main() {

	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go send(ch1, 1)
	go send(ch2, 2)

	for {
		select {
		case e := <-ch1:
			fmt.Printf("rece from ch1:%d\n", e)
		case e := <-ch2:
			fmt.Printf("rece from ch2:%d\n", e)
		default:
			fmt.Printf("no element ffrom ch1 and ch2.\n")
			time.Sleep(time.Second)
		}
	}
}
