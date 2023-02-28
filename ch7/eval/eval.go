package eval

import (
	"fmt"
	"math"
)

type Var string

type literal float64

type Env map[Var]float64

type Expr interface {
	Eval(env Env) float64
}

// unary 一元操作符及其操作数
type unary struct {
	op rune
	x  Expr
}

// unary 二元操作符及其操作数
type binary struct {
	op   rune
	x, y Expr
}

type call struct {
	fn   string
	args []Expr
}

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(env Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator : %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsurppoted binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsurppored function call: %s", c.fn))
}
