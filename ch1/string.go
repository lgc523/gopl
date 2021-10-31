package main

import (
	"bytes"
	"fmt"
	"path"
	"path/filepath"
	"strconv"
	strings2 "strings"
	"unicode"
)

func main() {
	//strings()
	//bytesTest()
	//unicodeTest()
	//pathTest()
	strconvTest()
}

// bytes2 slice操作
func bytesTest() {
	s1 := "hello bytes"
	s2 := "hello bytes"
	contains := bytes.Contains([]byte(s1), []byte(s2))
	fmt.Println(contains)
	var buf bytes.Buffer
	buf.WriteString("hello")
	buf.WriteByte(128)
	buf.WriteRune('√')
	fmt.Println(buf.String()) //hello�√
}

func strings() {
	s1 := "hello strings"
	s2 := "hello strings"
	compare := strings2.Compare(s1, s2)
	fmt.Println(compare)
}

func unicodeTest() {
	digit := unicode.IsDigit('a')
	fmt.Println(digit)
}

func pathTest() {
	dir := path.Dir("~/spider/Desktop/a.txt")
	fmt.Println(dir)
	abs, _ := filepath.Abs("~/spider/Desktop/a.txt")
	fmt.Println(abs)
}

func strconvTest(){
	//integer to ascii
	is := strconv.Itoa(1)
	fmt.Println(is)
	fmt.Println(fmt.Sprintf("%d",1))
	fmt.Println(strconv.FormatInt(523,2))
	in, _ := strconv.Atoi("523")
	fmt.Println(in)
	parseInt, _ := strconv.ParseInt("-523", 10, 11)
	fmt.Println(parseInt)
	parseUint, _ := strconv.ParseUint("523", 10, 11)
	fmt.Println(parseUint)
}
