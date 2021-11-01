package main

import (
	"fmt"
	"math"
)

var global *int

func main() {
	fmt.Println(-5 % 3)  //-2
	fmt.Println(-5 % -3) //-2
	fmt.Println(+1)      //1
	fmt.Println(-1)      //-1
	fmt.Println(-1.1)    //-1.1
	fmt.Println(+1.1)    //1.1
	var u uint8 = 255
	//255 0 1
	fmt.Println(u, u+1, u*u)
	var i int8 = 127
	fmt.Println(i, i+1, i*i)
	//127 -128 1
	fmt.Println(math.NaN())
	//NaN
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)
	//0 -0 +Inf -Inf NaN
	fmt.Println(math.IsNaN(z / z)) //true
	s := "hello"
	t := s
	s += ",golang"
	fmt.Println(s)    //hello,golang
	fmt.Println(t)    //hello
	fmt.Println(s[0]) //104
	fmt.Println(`原生字符串可以这么写,转义序列不起作用,可以展开，用在HTML模版、JSON字面量、命令行提示、正则表达式`)
	var st = `\377`
	var ss = `\x377`
	fmt.Printf("%o\n", st)
	fmt.Printf("%x\n", ss)
	fmt.Printf("%x", `\u4e16`)
	fmt.Println('世' == '\u4e16')
	fmt.Println(string('\uFFFD'))
}
func f() {
	var x int
	x = 1
	// f() 返回 global 还可以访问 x ，x 从 f 中逃逸，x 使用堆空间
	global = &x
}

func g() {
	y := new(int)
	*y = 1
}
