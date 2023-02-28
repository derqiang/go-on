package main

import (
	"fmt"
	"go_try/ch6/geometry"
	"go_try/ch6/intset"
	"go_try/me"
	"image/color"
	"sync"
)

func main() {
	//TryAndTry()
	//

	ch6Runner := []me.Runner{
		intset.ISRunner(1),
	}
	ch6Runner[0].Run()
}

func TryAndTry() {

	p := geometry.Point{1, 2}
	q := geometry.Point{3, 4}
	fmt.Println(geometry.Distance(p, q))
	fmt.Println(p.Distance(q))

	perim := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
	fmt.Println(geometry.Path.Distance(perim))

	r := &geometry.Point{1, 2}
	r.Distance(geometry.Point{1, 2})
	r.ScaleBy(2)
	fmt.Printf("\nScaled Point : %#v\n", *r)

	il := &IntList{
		1, &IntList{
			2,
			&IntList{
				3,
				nil,
			},
		},
	}
	println(il.Sum())

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var pp = geometry.ColorPoint{geometry.Point{1, 1}, red}
	var qq = geometry.ColorPoint{geometry.Point{5, 4}, blue}
	fmt.Println(pp.Distance(qq.Point))
}

type IntList struct {
	Value int
	Tail  *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	//sum := 0
	//for next := list; next != nil; next = next.Tail {
	//	fmt.Printf("processing .. %#v\n", next)
	//	sum += next.Value
	//}
	return list.Value + list.Tail.Sum()
}

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock()
	defer cache.Unlock()
	return cache.mapping[key]
}
