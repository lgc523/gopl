package main

import (
	"fmt"
	"os"
)

func main() {
	var buf [512]byte
	read, err := os.Stdin.Read(buf[:])
	if err != nil {
		fmt.Println("read error:", err)
		return
	}
	fmt.Println("count:", read, ",msg:", string(buf[:]))
}
