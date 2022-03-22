package main

import "sync"

func main() {
	raceRange()
}

func rangeSafety() {
	wg := sync.WaitGroup{}

	sli := []int{1, 2, 3, 4, 5}

	for i := range sli {
		wg.Add(1)
		go func(temp int) {
			println(temp)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func raceRange() {
	wg := sync.WaitGroup{}

	sli := []int{1, 2, 3, 4, 5}

	for i := range sli {
		wg.Add(1)
		//share i
		go func() {
			println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}
