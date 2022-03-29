package geometry

// References:
// 1. https://gobyexample.com/interfaces

type Geometry interface {
	Name() string
	Volume() float64
	Surface() float64
	Perim() float64
}
