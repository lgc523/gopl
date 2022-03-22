package dsa

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) *tree {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	ints := appendValues(values[:0], root)
	fmt.Printf("sort append slice:%v\n", ints)
	return root
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, v int) *tree {
	if t == nil {
		t := new(tree)
		t.value = v
		return t
	}
	if v < t.value {
		t.left = add(t.left, v)
	} else {
		t.right = add(t.right, v)
	}
	return t
}
func printTree(root *tree) {
	if root != nil {
		fmt.Println(root.value)
		printTree(root.left)
		printTree(root.right)
	}
}
