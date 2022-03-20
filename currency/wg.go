package main

import (
	"fmt"
	"sync"
)

const (
	noGoroutine = 5
	noTask      = 10
)

var wg sync.WaitGroup

func main() {
	tasks := make(chan int, noTask)
	for no := 1; no <= noGoroutine; no++ {
		wg.Add(1)
		go taskProcessor(tasks, no)
	}

	for taskNo := 1; taskNo <= noTask; taskNo++ {
		tasks <- taskNo
	}

	wg.Wait()
	close(tasks)
}

func taskProcessor(tasks chan int, workerNo int) {
	defer wg.Done()
	for t := range tasks {
		fmt.Printf("Worker %d is processing Task no: %d \n", workerNo, t)
	}
	fmt.Printf("Worker %d got off work \n", workerNo)
}
