package ch1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func BenchTest() {
	for i := 0; i < 20; i++ {
		println(" === " + strconv.Itoa(i) + " === ")
		s := time.Now()

		for i := 0; i < 100000; i++ {
			echo2()
		}
		fmt.Printf(" ECHO2 : %v \n", time.Now().Sub(s))

		s = time.Now()
		for i := 0; i < 100000; i++ {
			echo3()
		}
		fmt.Printf(" ECHO3 : %v \n", time.Now().Sub(s))
	}
}

func echo3() {

	strings.Join(os.Args[1:], " ")
	//fmt.Println(os.Args[1:])
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	//fmt.Println(s)
}

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func main1() {

	str := "aaaaaaaaAaaB我我我我我我我我我我我我我我我我我我我我我们"
	runes := []rune(str)
	runes = append(runes, []rune("\\0")[0])
	bound := 0
	max := 0
	recorder := make([]int, 0)
	for i, _ := range runes {
		if strings.ToLower(string(runes[i])) != strings.ToLower(string(runes[bound])) {
			fmt.Printf(" bound: %v , cur: %v , span : %v \n", bound, i, string(runes[bound:i]))
			l := i - bound
			if max < l {
				max = l
				recorder = []int{i}
			} else if max == l {
				recorder = append(recorder, i)
			}
			bound = i
		}
	}
	fmt.Printf(" end list : %v , max : %v \n", recorder, max)
	for _, e := range recorder {
		fmt.Printf(" range: [%v,%v) , max : %v > %v \n", e-max, e, max, string(runes[e-max:e]))
	}

}
