package main

import "fmt"

type Point struct {
	X, y int
}

type Circle struct {
	Point
	//Center Point
	Radius int
}

func (c Circle) print() {
	fmt.Println(c.Radius + c.X)
}

type Wheel struct {
	Circle
	//Circle Circle
	Spokes int
}

func main() {
	var w Wheel
	//匿名结构体默认是类型名，.可选
	//w.Circle.Center.X
	w.X = 8
	w.Radius = 8
	fmt.Printf("%#v\n", w)
	w.print()
}
