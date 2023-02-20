package _defer

import "fmt"

type DeferRunner int64

func (df DeferRunner) Run() {
	//i := fn1()
	//i := fn2()
	fmt.Printf("Run Body 1: return = %d, g = %d\n", fn1(), g)
	fmt.Printf("Run Body 2: return = %d, g = %d\n", fn2(), g)
}

var g = 100

func fn1() (r int) {
	defer func() {
		g = 200
	}()
	fmt.Printf("f1: g = %d\n", g)
	return g
}

func fn2() (r int) {
	r = g
	defer func() {
		r = 200
	}()
	fmt.Printf("f2: g = %d\n", g)
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
