package main

import (
	"fmt"
	"go_try/ch9/banks/bank"
	"time"
)

func main() {
	b := bank.Bank1(1)
	fmt.Println("=", b.Balance())

	go func() {
		b.Deposit(200)
		fmt.Println("=", b.Balance())
	}()

	go b.Deposit(100)

	time.Sleep(10 * time.Second)

}

// 产生数据竞争的情况有三大类，A先B后，B先A后，AB各自原子步骤交错执行，一般也是交错执行导致最终结果不确定
func sliceRacer() {
	var x []int
	go func() { x = make([]int, 100) }()
	go func() { x = make([]int, 100000) }()
	x[99999] = 1
}
