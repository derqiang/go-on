package main

import (
	"fmt"
	"go_try/ch2/popcount"
)

func main() {
	println("==== RESULT ====")
	println(popcount.PopCount(5))

	// 258 * 8
	fmt.Printf("64 >> 0 : %.64b\n64 >> 8 : %.64b \n", 64>>0, 256>>(8*2))

	// 字符串值不可变
	s := "left foot \a \b \v"
	t := s
	s += ", right foot"
	fmt.Printf("%v > %v \n", t, s)

	s = "Hello, 世界"
	fmt.Printf(" ===> % x", s)
}
