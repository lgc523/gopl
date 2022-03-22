package main

import "fmt"

func rangeSliceTrap() {
	var sli = []int{1, 2, 3, 4, 5}
	for _, i := range sli {
		//range slice copy element
		i = i + 1
		fmt.Println(i)
	}
	for _, i := range sli {
		fmt.Println(i)
	}
	fmt.Println("---")
}

func main() {
	rangeSliceTrap()
	generalRangeModify()
}

func generalRangeModify() {
	var sli = []int{1, 2, 3, 4, 5}
	for i, v := range sli {
		sli[i] = v + 1
	}
	for _, i := range sli {
		fmt.Println(i)
	}
}
