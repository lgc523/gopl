package main

import "fmt"

func main() {

	ch := make(chan int, 3)
	ch <- 2
	ch <- 3
	close(ch)
	for e := range ch {
		fmt.Printf("get element from ch:%d\n", e)
	}

}
