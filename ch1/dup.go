package ch1

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ChapterFirstToThreeDup1() {

	// Chapter 1.3, P6
	counts := make(map[string]int, 0)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if len(input.Text()) > 0 {
			counts[input.Text()]++
			continue
		}
		break
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// ChapterDup2 Chapter 1.3 , 30, read multiple line from the existed file
func ChapterDup2() {
	counts := make(map[string]int)

	args := os.Args[1:]
	if len(args) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range args {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2 : %v \n", err)
				continue
			}
			countLines(f, counts)
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf(" %d \t %s \n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		counts[scan.Text()]++
	}
}

func ChapterDup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Dup3 : %v \n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf(" %d \t %s \n", n, line)
		}
	}
}
