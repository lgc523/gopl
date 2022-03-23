package main

import "fmt"

func main() {
	f := add()
	fmt.Println(f(10))
	fmt.Println(f(20))
	ff := add()
	fmt.Println(ff(10))
	fmt.Println(ff(20))

	b := add0(1)
	fmt.Println(b(10))
	fmt.Println(b(20))

	bb := add0(2)
	fmt.Println(bb(10))
	fmt.Println(bb(20))
}

func add() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func add0(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}
