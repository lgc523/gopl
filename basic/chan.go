package main

import (
	"fmt"
	"time"
)

func main() {
	//var c chan int
	//close(c)
	//fmt.Println(c == nil)
	ch := make(chan interface{})

	go func() {
		for {
			ch <- 0
			time.Sleep(time.Second)
		}
	}()
	readChan(ch)
}
func readChan(ch chan interface{}) {
	for {
		select {
		case out, _ := <-ch:
			fmt.Printf("rece:%v\n", out)
		default:
			fmt.Println("empty chan buf")
			time.Sleep(time.Second)
		}
	}
}
