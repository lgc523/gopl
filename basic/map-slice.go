package main

import (
	"fmt"
)

type Person struct {
	Name  string
	Likes []string
}

var people []*Person

func main() {
	likes := make(map[string][]*Person)
	for _, p := range people {
		for _, l := range p.Likes {
			likes[l] = append(likes[l], p)
		}
	}
	fmt.Printf("%d people like to  read & code, his/her name is %s.\n", len(likes["meinv"]), likes["meinv"][0].Name)
}

func init() {
	var p = Person{"spider", []string{"meinv"}}
	people = append(people, &p)
}
