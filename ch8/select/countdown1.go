package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//countDown1()
	//countDown2()
	//launchRocket()
	skipSelector()
}

func countDown1() {
	//tick := time.Tick(1 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	abort := countDown2()
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-ticker.C:
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	ticker.Stop()
	println("END!")
}

func countDown2() chan struct{} {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	return abort
}

func launchRocket() {
	abort := countDown2()
	select {
	case <-time.After(10 * time.Second):
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
}

func skipSelector() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			println(x)
		case ch <- i:
		}
	}
}
