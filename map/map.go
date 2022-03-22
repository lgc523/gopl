package main

import "fmt"

func main() {
	var nm map[string]int
	fmt.Println(nm == nil)
	m := make(map[string]int)
	fmt.Println(m == nil)
	fmt.Printf("init\n map %v,len:%d,addr:%p\n", m, len(m), &m)
	add(m)
	fmt.Printf("after add\n map %v,len:%d,addr:%p\n", m, len(m), &m)
	for v := range m {
		fmt.Println(v)
	}
	changeMap(m)
	fmt.Printf("after changeMap\n map %v,len:%d,addr:%p\n", m, len(m), &m)
	for v := range m {
		fmt.Println(v)
	}
	mapPoint(&m)
}

func add(m map[string]int) {
	m["a"] = 1
	m["d"] = 1
	m["add"] = 1
	m["adda"] = 1
	fmt.Printf("add\n map %v,len:%d,addr:%p\n", m, len(m), m)
}

func changeMap(m map[string]int) {
	m = nil
	fmt.Printf("changeMap\n map %v,len:%d,addr:%p\n", m, len(m), m)
}

func mapPoint(m *map[string]int) {
	fmt.Printf("mapPoint:m %p", m)
}
