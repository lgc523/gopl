package main

import "fmt"

type Foo struct {
	Name string
	age  int
}

func main() {
	p := &Foo{"spider", 25}
	q := &Foo{"spider", 25}
	fmt.Println(q.Name == q.Name && p.age == q.age)
	fmt.Println(p == q)
}
