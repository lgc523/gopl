package dsa

import (
	"testing"
)

func TestTree(t *testing.T) {

	s := []int{1, 8, 5, 3, 8, 0, 5}
	st := Sort(s)
	printTree(st)
}
