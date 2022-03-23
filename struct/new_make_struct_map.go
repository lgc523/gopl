package main

type Alpha map[string]string
type Bar struct {
	one string
	two int
}

func main() {
	a := new(Bar)
	a.one = "a"
	(*a).one = "one"

	//compiler error
	//b :=make(Bar)

	alpha := make(Alpha)
	alpha["alpha"] = "alpha"

	c := new(Alpha)
	//panic: assignment to entry in nil map
	(*c)["c"] = "c"
}
