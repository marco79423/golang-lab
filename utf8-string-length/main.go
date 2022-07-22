package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(utf8.RuneCountInString("abc")) // 1-byte 字元 -> Output: 3
	fmt.Println(utf8.RuneCountInString("中文"))  // 3-byte 字元 -> Output: 2
	fmt.Println(utf8.RuneCountInString("𒀀𒀀𒀀")) // 4-byte 字元 -> Output: 3
	fmt.Println(utf8.RuneCountInString("a中𒀀")) // 混合 -> Output: 3
}
