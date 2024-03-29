package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch1 := make(chan int)
	go func() {
		fmt.Printf("\tbackend start...\n")
		//time.Sleep(5 * time.Second)
		for in := range ch1 {
			fmt.Printf("\tbackend goroutine : %v\n", in)
		}
		//fmt.Printf("\tbackend goroutine : %v\n", <-ch1)
		//fmt.Printf("\t[backend end]! - %v\n", <-ch1)
		fmt.Printf("\tbackend end!\n")
	}()

	//ch1 <- 1
	//fmt.Printf("main get 1...%v\n", <-ch1)
	//ch1 <- 2
	//fmt.Printf("main get 2...%v\n", <-ch1)
	//ch1 <- 3
	//fmt.Printf("main get 3...%v\n", <-ch1)
	//ch1 <- 4
	//fmt.Printf("main get 4...%v\n", <-ch1)
	fmt.Printf("main end!\n")
	//cakeRoom()
	time.Sleep(1000 * time.Second)
}

func cakeRoom() {
	var wg sync.WaitGroup

	belt := []chan cakeUnit{
		make(chan cakeUnit),
		make(chan cakeUnit),
	}

	workers := []*worker{
		{"worker 3", IDLE, belt[1], nil, "\t\t"},
		{"worker 2", IDLE, belt[0], belt[1], "\t"},
		{"worker 1", IDLE, nil, belt[0], ""},
	}

	for _, w := range workers {
		var ww = w
		ww.handleCake(&wg)
	}

	wg.Wait()
	for _, w := range workers {
		if w.out != nil {
			close(w.out)
			fmt.Printf("clean : len(w.out) = %v\n", len(w.out))
		}
	}
}

const (
	IDLE = iota + 1
	DOING
)

type cakeUnit struct {
	name string
	id   int
}

func (cu cakeUnit) String() string {
	return fmt.Sprintf("{id: %d, name : %v}", cu.id, cu.name)
}

type worker struct {
	name  string
	state int
	in    <-chan cakeUnit
	out   chan<- cakeUnit
	flag  string
}

func (w *worker) handleCake(wg *sync.WaitGroup) {
	if w.in != nil {
		go func() {
			fmt.Printf("%vReady?  %v\n", w.flag, w.name)
			for in := range w.in {
				fmt.Printf("%vworking for : %v\n", w.flag, in)
				//var elapsed = time.Duration(rand.Int()%10 + 1)
				//time.Sleep(elapsed * time.Second)
				in.name = in.name + "-" + w.name
				wg.Done()
				go func(inner cakeUnit) { w.out <- inner }(in)
			}
		}()
	}

	if w.in == nil && w.out != nil {
		wg.Add(20)
		for c := 0; c < 10; c++ {
			todo := cakeUnit{
				name: w.name,
				id:   c,
			}
			fmt.Printf("%vworking for : %v\n", w.flag, todo)
			w.out <- todo
		}
	}
}
