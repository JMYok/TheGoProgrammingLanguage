package eval

import (
	"fmt"
	"math"
)

type Expr interface {
	// Eval返回表达式在env上下文下的值
	Eval(env Env) float64
	// Check方法报告表达式中的错误，并把表达式中的变量加入Vars中
	Check(vars map[Var]bool) error
}

// Var表示一个变量，比如x
type Var string

// 上下文：映射变量到值
type Env map[Var]float64

// literal是一个数字常量，比如3.141
type literal float64

// unary表示一元操作符表达式，比如-x
type unary struct {
	op rune //'+','-'中的一个
	x  Expr
}

// binary表示二元操作符表达式，比如x+y
type binary struct {
	op   rune //'+','-','*','/'中的一个
	x, y Expr
}

// call表示函数调用表达式，比如sin(x)
type call struct {
	fn   string // one of "pow","sin","sqrt"中的一个
	args []Expr
}

func (v Var) Eval(env Env) float64 {
	return env[v]
}
func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator:%q", u.op))
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
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
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
	panic(fmt.Sprintf("unsupported function call:%q", c.fn))
}
