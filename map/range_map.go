package main

import "fmt"

func rangeMapTrap() {
	sli := []int{1, 2, 3, 4, 5}
	m := make(map[int]*int)
	for k, v := range sli {
		//v addr is copy out
		m[k] = &v
	}
	for k, v := range m {
		fmt.Printf("k:%v,v:%v\n", k, *v)
	}
	fmt.Println(m)
}

func rangeMap() {
	sli := []int{1, 2, 3, 4, 5}
	m := make(map[int]*int)
	for k, v := range sli {
		p := v
		m[k] = &p
	}
	for k, v := range m {
		fmt.Printf("k:%v,v:%v\n", k, *v)
	}
	fmt.Println(m)
	fmt.Println("---")
}

func main() {
	rangeMapTrap()
	rangeMap()
}
