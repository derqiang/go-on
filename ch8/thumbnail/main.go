package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	//input := bufio.NewScanner(os.Stdin)
	//for input.Scan() {
	//	println(input.Text())
	//	println(input.Text())
	//	thumb, err := ImageFile(input.Text())
	//	if err != nil {
	//		log.Print(err)
	//		continue
	//	}
	//	fmt.Println(thumb)
	//}
	//if err := input.Err(); err != nil {
	//	log.Fatal(err)
	//}

	filenames := make(chan string, 10)
	go func() {
		for i := 0; i < 7; i++ {
			path := fmt.Sprintf("./ch8/thumbnail/assets/%d.jpg", i)
			filenames <- path
		}
		close(filenames)
	}()
	log.Println(makeThumbnail6(filenames))
	log.Println(">>>> END >>>>")
}

/// 等待和协调子任务执行结果
func makeThumbnail3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			ImageFile(f)
			ch <- struct{}{}
		}(f)
	}

	for range filenames {
		<-ch
	}
}

/// 如果结果执行过程出现错误，咋办? 一个简单的想法是，不管是否有错误，都返回错误（没有错误则返回nil了 ）
// 隐藏bug：当第一个非nil的错误发生时，[errs]这个通道，因为函数执行完毕，导致通道并没有释放资源，而出现goroutine泄露
func makeThumbnail4(filenames []string) error {
	errs := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := ImageFile(f)
			errs <- err
		}(f)
	}

	for range filenames {
		if err := <-errs; err != nil {
			return err
		}
	}

	return nil
}

/// 上述隐藏的问题，另外增加一个需求：将成功执行的结果进行返回
func makeThumbnail5(filenames []string) (thumbs []string, err error) {
	defer traceRun("makeThumbnail5")()
	type item struct {
		thumb string
		err   error
	}
	ch := make(chan item, len(filenames))
	for i, f := range filenames {
		go func(f string, i int) {
			var it item
			if i%2 == 1 {
				time.Sleep(time.Duration(i) * time.Second)
				it.thumb, it.err = f, fmt.Errorf("err : %v", i)
			} else {
				it.thumb, it.err = f, nil
			}
			ch <- it
			println("!GO ROUTINE OVER!")
		}(f, i)
	}
	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbs = append(thumbs, it.thumb)
	}
	println("!OVER!")
	return thumbs, nil
}

func makeThumbnail6(filenames <-chan string) int64 {

	var wg sync.WaitGroup
	sizes := make(chan int64)
	for f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer func() {
				wg.Done()
			}()
			out, err := ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			st, _ := os.Stat(out)
			sizes <- st.Size()
		}(f)
		println("1 ---- IN ----")
	}

	go func() {
		wg.Wait()
		close(sizes)
		println("Close sizes channel.")
	}()

	var total int64
	for s := range sizes {
		fmt.Printf("....>>> main goroutine runnming : %v\n", s)
		total += s
	}

	return total
}

func traceRun(msg string) func() {
	now := time.Now()
	log.Printf("enter for [%v], at: %v", msg, now)
	return func() {
		log.Printf("exit for [%v] at : %v", msg, time.Now())
	}
}
