package intset

import (
	"bytes"
	"fmt"
)

// 数据密度

type ISRunner int64
type IntSet struct {
	words []uint64
}

func (s ISRunner) Run() {
	// 1000000010
	// 514
	println(20 << (1))
	var t IntSet
	for i := 0; i < 128; i++ {
		t.Add(i)
	}
	fmt.Println(t.String())
	fmt.Println(t.words)

	fmt.Printf("%d\n", 1<<9|1<<1)
	var x, y IntSet
	x.Add(1)
	//x.Add(144)
	x.Add(9)
	fmt.Println(x.String())
	fmt.Println(x.words)

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())
	fmt.Println(x.words)

	x.UnionWith(&y)
	fmt.Println(x.String())
	fmt.Println(x.Has(9), x.Has(123))
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (*IntSet) Len() int {
	return 0
}

func (*IntSet) Remove(x int) int {
	return 0
}

func (*IntSet) Clear() int {
	return 0
}

func (*IntSet) Copy() *IntSet {
	return &IntSet{}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
