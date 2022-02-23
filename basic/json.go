package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	Age   int    `json:"age"`
	Name  string `json:"name"`
	Child []int  `json:"child"`
}

func main() {
	json0 := `{"age":25, "name":"spider", "child":[1,2,3]}`
	f0 := Foo{}
	json.Unmarshal([]byte(json0), &f0)
	c := f0.Child
	fmt.Println(c)

	json1 := `{"age":"18","name":"spider","child":[3,4,5,6,7,8]}`
	json.Unmarshal([]byte(json1), &f0)
	fmt.Println(c)
}
