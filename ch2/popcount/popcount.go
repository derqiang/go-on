package popcount

import "fmt"

var pc [256]byte

func init() {
	count := 0
	for i := range pc {
		count++
		if i%8 == 0 {
			println()
		}
		pc[i] = pc[i/2] + byte(i&1)
		fmt.Printf("%v ", pc[i])
	}
	fmt.Printf("\ncount : %v \n", count)
}

func init() {
	fmt.Println("Second Init")
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
