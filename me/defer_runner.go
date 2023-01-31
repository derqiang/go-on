package me

import "fmt"

var g = 100

type DeferRunner int64

func (df DeferRunner) Run() {
	//i := fn1()
	i := fn2()
	fmt.Printf("Run Body: i = %d, g = %d\n", i, g)
}

func fn1() (r int) {
	defer func() {
		g = 200
	}()
	fmt.Printf("f: g = %d\n", g)
	return g
}

func fn2() (r int) {
	r = g
	defer func() {
		r = 200
	}()
	fmt.Printf("f: g = %d\n", g)
	r = 0
	return r
}

///  延迟+recover捕获崩溃异常
func panicExecute() {
	fmt.Println("Start to go ... ")
	defer func() {
		fmt.Println("Defer Come in.")
		if err := recover(); err != nil {
			fmt.Printf("An Error Occurred! : %v", err)
		}
	}()
	panic("An Panic Occurred")
	fmt.Println("End Panic Execute!")
}
