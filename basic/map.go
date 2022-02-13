package main

import "fmt"

func main() {

}

func init() {
	//create map

	m := map[string]int{
		"beijing":  1,
		"niaochao": 2,
	}
	for k, v := range m {
		fmt.Printf("%s-%d\n", k, v)

	}
	init1()
}

func init1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover in init..%v", r)
		}
	}()
	var m map[string]int
	fmt.Printf("m is nil?%v\n", m == nil)
	val := m["nil"]
	fmt.Println(val)
	m["nil"] = -1
}
