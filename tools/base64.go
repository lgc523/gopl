package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

//尝试写一个base64工具

var base64Map map[uint8]rune

var param = flag.String("param", "", "trans your code use base64 codeMap")

func main() {
	flag.Parse()
	//第一个参数 ~ 解析成了用户目录，需要加 ''
	str := os.Args[1:2][0]
	runeSize := utf8.RuneCountInString(str)
	fmt.Printf("input string is: [%s],codePoint len:[%d],byte:[%d]\n", str, runeSize, runeSize)
	for i := 0; i < runeSize; i++ {
		codePoint, _ := utf8.DecodeRuneInString(str[i : i+1])
		asc := strconv.QuoteToASCII(string(codePoint))
		r := []rune(asc)
		fmt.Printf("codePoint decimal:[%v]->ascii[%s]->rune:%v->binary:[%b]\n", codePoint, asc, r, r)
	}
	//fmt.Printf("base64 is: %s\n",str)

}

// 建立编码表mapping
func init() {
	base64Map = make(map[uint8]rune)
	base64Map[0] = 'A'
	base64Map[1] = 'B'
	base64Map[2] = 'C'
	base64Map[3] = 'D'
	base64Map[4] = 'E'
	base64Map[5] = 'F'
	base64Map[6] = 'G'
	base64Map[7] = 'H'
	base64Map[8] = 'I'
	base64Map[9] = 'J'
	base64Map[10] = 'K'
	base64Map[11] = 'L'
	base64Map[12] = 'M'
	base64Map[13] = 'N'
	base64Map[14] = 'O'
	base64Map[15] = 'P'
	base64Map[16] = 'Q'
	base64Map[17] = 'R'
	base64Map[18] = 'S'
	base64Map[19] = 'T'
	base64Map[20] = 'U'
	base64Map[21] = 'V'
	base64Map[22] = 'W'
	base64Map[23] = 'X'
	base64Map[24] = 'Y'
	base64Map[25] = 'Z'

	base64Map[26] = 'a'
	base64Map[27] = 'b'
	base64Map[28] = 'c'
	base64Map[29] = 'd'
	base64Map[30] = 'e'
	base64Map[31] = 'f'
	base64Map[32] = 'g'
	base64Map[33] = 'h'
	base64Map[34] = 'i'
	base64Map[35] = 'j'
	base64Map[36] = 'k'
	base64Map[37] = 'l'
	base64Map[38] = 'm'
	base64Map[39] = 'n'
	base64Map[40] = 'o'
	base64Map[41] = 'p'
	base64Map[42] = 'q'
	base64Map[43] = 'r'
	base64Map[44] = 's'
	base64Map[45] = 't'
	base64Map[46] = 'u'
	base64Map[47] = 'v'
	base64Map[48] = 'w'
	base64Map[49] = 'x'
	base64Map[50] = 'y'
	base64Map[51] = 'z'

	base64Map[52] = '0'
	base64Map[53] = '1'
	base64Map[54] = '2'
	base64Map[55] = '3'
	base64Map[56] = '4'
	base64Map[57] = '5'
	base64Map[58] = '6'
	base64Map[59] = '7'
	base64Map[60] = '8'
	base64Map[61] = '9'

	base64Map[62] = '+'
	base64Map[63] = '/'
}
