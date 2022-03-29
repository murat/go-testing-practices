package math

// TODO: change test file package name

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"happy path", args{3, 5}, 8},
		// {"unhappy path", args{3, 5}, 7}, // cannot be succeed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSub(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"happy path#positive", args{5, 3}, 2},
		{"happy path#negative", args{3, 5}, -2},
		// {"unhappy path", args{3, 5}, 7}, // cannot be succeed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sub(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"happy path", args{3, 5}, 15},
		// {"unhappy path", args{3, 5}, 14}, // cannot be succeed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiply(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"happy path", args{10, 5}, 2},
		{"happy path", args{10, 5}, 2.0},
		// {"unhappy path", args{10, 5}, 3}, // cannot be succeed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Divide(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEven(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"happy path", args{10}, true},
		{"happy path", args{5}, false},
		// {"unhappy path", args{10}, false}, // cannot be succeed
		// {"unhappy path", args{5}, true},   // cannot be succeed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Even(tt.args.a); got != tt.want {
				t.Errorf("Even() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOdd(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"happy path", args{10}, false},
		{"happy path", args{5}, true},
		// 	{"unhappy path", args{10}, true}, // cannot be succeed
		// 	{"unhappy path", args{5}, false}, // cannot be succeed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Odd(tt.args.a); got != tt.want {
				t.Errorf("Odd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoot(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"happy path", args{4}, 2},
		{"happy path", args{8}, 2.82842712474619},
		{"happy path", args{1}, 1},
		// {"unhappy path", args{4}, 3}, // cannot be succeed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Root(tt.args.a); got != tt.want {
				t.Errorf("Root() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPow(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"happy path", args{2, 2}, 4},
		{"happy path", args{1, 100}, 1},
		// {"unhappy path", args{2, 3}, 10}, // cannot be succeed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pow(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Pow() = %v, want %v", got, tt.want)
			}
		})
	}
}
