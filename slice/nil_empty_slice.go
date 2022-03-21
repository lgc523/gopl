package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//底层数组没有分配
	var a []int
	//底层数组指针非空
	b := make([]int, 0)

	//see makeslice mallocgc zerobase

	if a == nil {
		fmt.Println("a is nil")
	} else {
		fmt.Println("a is not nil")
	}
	if b == nil {
		fmt.Println("b is nil")
	} else {
		fmt.Println("b is not nil")
	}

	as := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	bs := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	fmt.Printf("len=%d,cap=%d,type=%d\n", len(a), cap(a), as.Data)
	fmt.Printf("len=%d,cap=%d,type=%d\n", len(b), cap(b), bs.Data)
}
