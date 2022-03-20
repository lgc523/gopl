package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	RMB
)

func main() {
	symbol := [...]string{USD: "$", RMB: "¥", EUR: "¢"}
	for i, v := range symbol {
		fmt.Println(i, v)
	}
	fmt.Println(RMB, symbol[RMB])
}
