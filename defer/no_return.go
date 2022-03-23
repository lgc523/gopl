package main

import "fmt"

func main() {

	fmt.Println(returnN())
}

func returnN() (result string) {
	defer func() {
		//外部 panic
		//recover 终止 获取到值
		if p := recover(); p != nil {
			//type select
			r, ok := p.(string)
			if ok {
				result = r
			} else {
				panic("p type select err")
			}
		}
	}()
	panic("panic")
}
