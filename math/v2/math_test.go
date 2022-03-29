package v2_test

import (
	"testing"

	math "github.com/murat/testing-practices/math/v2"
)

func TestAdd(t *testing.T) {
	t.Run("happy path via integers", func(t *testing.T) {
		if got := math.Add(3, 5); got != 8 {
			t.Errorf("Add(3, 5) = 8 but %v", got)
		}
	})
	t.Run("happy path via float64", func(t *testing.T) {
		if got := math.Add(3.5, 5.5); got != 9 {
			t.Errorf("Add(3, 5) = 8 but %v", got)
		}
	})

	// following test never succees
	// t.Run("unhappy path via integers", func(t *testing.T) {
	// 	if got := math.Add[int64](3, 5); got != 7 {
	// 		t.Errorf("Add(3, 5) = 7 but %v", got)
	// 	}
	// })

	// another way
	// manually comparing some functions outputs with some values (unnamed test cases)
	a := 1.2
	b := 2.3
	s := math.Add(a, b)
	if s != 3.5 {
		t.Errorf("1.2 + 2.3 should be 3.5, but it is %.2f", s)
	}
}

func FuzzAdd(f *testing.F) {
	f.Add(5, 10)
	f.Fuzz(func(t *testing.T, a int, b int) {
		r := math.Add(a, b)
		t.Log(a, b, r)
		if r != a+b {
			t.Errorf("a+b should be %v, but it is %v", a+b, r)
		}
	})
}

// func TestSub(t *testing.T) {
// 	type args struct {
// 		x T
// 		y T
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want T
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Sub(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Sub() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestMultiply(t *testing.T) {
// 	type args struct {
// 		x T
// 		y T
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want T
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Multiply(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Multiply() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestDivide(t *testing.T) {
// 	type args struct {
// 		x T
// 		y T
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want T
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Divide(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Divide() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestEven(t *testing.T) {
// 	type args struct {
// 		x T
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Even(tt.args.x); got != tt.want {
// 				t.Errorf("Even() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestOdd(t *testing.T) {
// 	type args struct {
// 		x T
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Odd(tt.args.x); got != tt.want {
// 				t.Errorf("Odd() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestRoot(t *testing.T) {
// 	type args struct {
// 		x T
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want float64
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Root(tt.args.x); got != tt.want {
// 				t.Errorf("Root() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestPow(t *testing.T) {
// 	type args struct {
// 		x T
// 		y T
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want T
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Pow(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Pow() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
