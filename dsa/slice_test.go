package dsa

import (
	"fmt"
	"testing"
)

func TestSlicePointParam(t *testing.T) {
	var s []int
	testM(&s)
	fmt.Printf("after slice:%v", s)

}

func testM(s *[]int) {
	*s = append(*s, 1)
	*s = append(*s, 2)
	*s = append(*s, 3)
	fmt.Printf("slice:%v", *s)
}
