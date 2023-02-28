package ch8

import (
	"fmt"
	"go_try/me"
	"time"
)

var Runners = []me.Runner{
	//GoroutineRunner(1),
	//ClockServer(0),
	//ClockServer(1),
	ClockServer(2),
}

type GoroutineRunner int64

func (GoroutineRunner) Run() {
	// 1 2 3 4 5 6 ...
	// 1 1 2 3 5 8
	go inProcess()
	fmt.Printf("\rFib(45) = %d\n", fib(45))
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func inProcess() {
	for {
		for _, c := range `-\|/` {
			fmt.Printf("\r%c", c)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
