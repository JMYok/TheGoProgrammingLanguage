package main

import (
	"errors"
	"testing"
)

type operate func(int, int) int

// 方案1
func calculate(x int, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

// 方案2
type calculateFunc func(int, int) (int, error)

func genCalculator(op operate) calculateFunc {
	return func(x int, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

func TestFuncAsParam(t *testing.T) {
	op := func(x int, y int) int {
		return x + y
	}
	t.Log(calculate(1, 2, op))
}

func TestFuncAsRes(t *testing.T) {
	x, y := 99, 100
	op := func(x int, y int) int {
		return x + y
	}

	add := genCalculator(op)
	result, err := add(x, y)
	t.Logf("The result: %d (error: %v)\n", result, err)
}
