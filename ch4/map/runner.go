package _map

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

type MapContainer int64

func (c MapContainer) Run() {
	//ages := make(map[string]int)
	//ages = map[string]int{
	//	"DQ": 2023 - 1989,
	//	"AA": 2023 - 1992,
	//	"MM": 2023 - 1991,
	//}
	//println("遍历Map ： ")
	//for k, v := range ages {
	//	println(k, v)
	//}
	//
	//println("按key排序取值： ")
	//var names = make([]string, 0, len(ages))
	//for k := range ages {
	//	names = append(names, k)
	//}
	//sort.Strings(names)
	//fmt.Printf("排序后的names: %v, \n", names)
	//
	//edges := map[int]Bag{
	//	0: {
	//		1: 0,
	//	},
	//	1: {
	//		2: 0, 3: 0, 4: 0,
	//	},
	//	2: {
	//		3: 0, 4: 0,
	//	},
	//}
	//
	//fmt.Printf("edges :  %v \n", edges)

	//CharCount()

}

type Bag map[int]int

// 同时统计输入源中含有的各个Unicode字符数量以及UTF8编码的占用字节对应的数量
func CharCount() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			println("EOF!!")
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 chracters \n", invalid)
	}
}

// 图结构的表达
var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
