package array

import (
	"crypto/sha256"
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}

type Runner int64

func (cur Runner) Run() {
	println(">>> Chapter 4.1 Array P_120 ~ P_123 >>>")
	println("StartIndex" + fmt.Sprintf("%v", cur))

	// 数组的初始化形式
	r := [...]int{99: -1}
	fmt.Println(r)

	symbol := [...]string{USD: "$", RMB: "¥", EUR: "€", GBP: "£"}
	fmt.Println(symbol)

	//
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("% x\n% x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	println("<<< Chapter 4.1 Array P_120 ~ P_123 <<<")
}
