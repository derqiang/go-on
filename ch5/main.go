package main

import (
	_defer "go_try/ch5/defer"
	"go_try/me"
	"math"
)

func add(x, y int) int {
	return x + y
}

// 减法函数
func sub(x int, y int) (z int) {
	z = x - y
	return
}

func main() {
	var _ = []me.Runner{
		FindLinkRunner(0),
		OutlineRunner(1),
		FindLinkRunner2(2),
		SquaresRunner(3),
		_defer.DeferRunner(4),
		Prereqs{},
		FindLinkRunner3(6),
		TitleRunner(7),
	}

	BigSlowOperation()
	//chapter5Runners[7].Run()

	_ = math.Sin(1)

	//if err := WaitForServer("https://daqiang.me"); err != nil {
	//	log.Fatalf("Site is donw: %v\n", err)
	//}
}
