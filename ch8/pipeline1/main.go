package main

import "time"

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	counter(ch1)
	squarer(ch1, ch2)

	for x := range ch2 {
		println(x)
	}
}

func counter(ch1 chan<- int) {

	go func() {
		for count := 0; count < 10; count++ {
			ch1 <- count
			time.Sleep(1 * time.Second)
		}
		close(ch1)
	}()
}

func squarer(ch1 <-chan int, ch2 chan<- int) {
	//go func() {
	//	for {
	//		x, ok := <-ch1
	//		if !ok {
	//			break
	//		}
	//		ch2 <- x * x
	//	}
	//}()
	go func() {
		for x := range ch1 {
			ch2 <- x * x
		}
		close(ch2)
	}()
}
