package main

import "fmt"

func main() {

	a := [5]int{1, 2, 3, 4, 5}
	b := a[1:4]
	fmt.Printf("a:%v\n", a)
	b = append(b, 0)

	c := a[1:3:3]
	fmt.Printf("c addr:%p\n", c)
	fmt.Printf("c:%v\n", c)
	c = append(c, 0)
	fmt.Printf("c addr:%p\n", c)
	fmt.Printf("a:%v\n", a)
	fmt.Printf("b:%v\n", b)
	fmt.Printf("c:%v\n", c)
}
