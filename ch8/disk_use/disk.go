package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")
var done = make(chan struct{})

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	var wg sync.WaitGroup
	fileSizes := make(chan int64)
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, fileSizes, &wg)
	}

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	var nfiles, nbytes int64
	var tick <-chan time.Time
	if *verbose {
		println(" VERBOSE ")
		tick = time.Tick(500 * time.Millisecond)
	}
loop:
	for {
		select {
		case <-done:
			for range fileSizes {
			}
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes, false)
		}
	}
	printDiskUsage(nfiles, nbytes, true)
	panic("test")
}

func canceled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func walkDir(dir string, fileSizes chan<- int64, n *sync.WaitGroup) {
	defer n.Done()
	if canceled() {
		return
	}
	for _, en := range dirents(dir) {
		if en.IsDir() {
			n.Add(1)
			subDir := filepath.Join(dir, en.Name())
			go walkDir(subDir, fileSizes, n)
		} else {
			fileSizes <- en.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}:
	case <-done:
		return nil
	}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64, ended bool) {
	fmt.Printf("%v%d files %.1f GB\n", func() string {
		if ended {
			return ""
		} else {
			return "\t\t"
		}
	}(), nfiles, float64(nbytes)/1e9)
}
