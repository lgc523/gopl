package main

import "fmt"

func main() {
	sli := []string{"s", "l", "i", "", "c", "e"}
	strings := nonempty(sli)
	for _, v := range strings {
		fmt.Println(v)
	}
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
