package main

type SquaresRunner int64

func (sr SquaresRunner) Run() {
	var f = square()
	println(f())
	println(f())
	println(f())
	println(f())
}

func square() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
