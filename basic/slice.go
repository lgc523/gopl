package main

import "fmt"

func main() {

	s := *new([]int)
	fmt.Println(s == nil)
	s = append(s, 2, 3, 5)
	fmt.Printf("slice len:%d,cap:%d\n", len(s), cap(s))
	for i, v := range s {
		fmt.Printf("index:%v - val:%v\n", i, v)
	}
	//copy
	ss := make([]int, 5)
	copy(s, ss)
	for _, v := range s {
		fmt.Println(v)

	}
}
