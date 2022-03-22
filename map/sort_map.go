package main

import (
	"fmt"
	"sort"
)

var m = map[string]int{
	"spider": 25, "nns": 25, "jasines": 24,
}

func sortMap(m map[string]int) {
	for k, v := range m {
		fmt.Printf("k:%v,v:%v\n", k, v)
	}
	keys := make([]string, len(m))
	i := 0
	for k, _ := range m {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	for _, v := range keys {
		fmt.Printf("k:%v,v:%v\n", v, m[v])
	}
}

func reverseMap(m map[string]int) {
	reverseM := make(map[int]string, len(m))

	for k, v := range m {
		reverseM[v] = k
	}
	for k, v := range reverseM {
		fmt.Printf("k:%v,v:%v\n", k, v)
	}
}

func main() {
	sortMap(m)
	reverseMap(m)
}
