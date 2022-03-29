package v2

import (
	"math"

	"golang.org/x/exp/constraints"
)

type Number interface {
	Int | Float
}
type Float interface {
	constraints.Float
}
type Int interface {
	constraints.Integer
}

func Add[T Number](x, y T) T {
	return x + y
}

func Sub[T Number](x, y T) T {
	return x - y
}

func Multiply[T Number](x, y T) T {
	return x * y
}

func Divide[T Number](x, y T) T {
	return x / y
}

// Even and Odd numbers are subsets of integers
// So, we are calling Even() and Odd() via only integers
func Even[T Int](x T) bool {
	return math.Mod(float64(x), 2) == 0
}

func Odd[T Int](x T) bool {
	return !Even(x)
}

func Root[T Number](x T) float64 {
	var sqt, root float64
	sqt = float64(x) / 2.0

	for {
		root = sqt
		sqt = (float64(x)/root + root) / 2
		if sqt == root {
			break
		}
	}

	return root
}

func Pow[T Number](x, y T) T {
	var r T
	r = 1
	for i := 0; i < int(y); i++ {
		r *= x
	}

	return r
}
