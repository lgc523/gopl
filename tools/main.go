package main

import (
	"fmt"
	"time"
)

func main(){
	a := []int{1,2,3}
	for i:= range a{
		a = append(a, i)
	}
	fmt.Println(a)
}

type A struct {
	B
	C map[string]D
}

type B struct {
	E, F  string
	G     string
	Timer H
}

type D struct {
	I uint64
}

type H struct {
	Timer time.Ticker
	J     chan D
}
