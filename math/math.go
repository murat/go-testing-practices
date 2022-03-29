package math

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b float64) float64 {
	return a / b
}

func Even(a int) bool {
	return a%2 == 0
}

func Odd(a int) bool {
	return !Even(a)
}

func Root(a int) float64 {
	var sqt, root float64
	sqt = float64(a) / 2

	for {
		root = float64(sqt)
		sqt = (float64(a)/root + root) / 2
		if sqt == root {
			break
		}
	}

	return root
}

func Pow(a, b int) int {
	r := 1
	for i := 0; i < b; i++ {
		r *= a
	}

	return r
}
