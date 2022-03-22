package main

import (
	"errors"
	"fmt"
	"os"
	"sync"
)

func main() {
	file, err := os.Create("record.txt")
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Error encounter: %w", err)
		}
		file.Close()
	}()
	if err != nil {
		panic(errors.New("Cannot create/open file"))
	}
	ss := []string{
		"james",
		"avery",
		"peter",
		"spider",
	}
	chanLock := make(chan int, 1)
	var wg sync.WaitGroup

	for _, str := range ss {
		wg.Add(1)
		go func(aString string) {
			chanLock <- 1
			for i := 0; i < 1000; i++ {
				file.WriteString(aString + "\n")
			}
			<-chanLock
			wg.Done()
		}(str)
	}
	wg.Wait()
}
