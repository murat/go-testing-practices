package main

import (
	"fmt"

	math "github.com/murat/testing-practices/math/v2"
)

func main() {
	// Math(v1)
	// fmt.Printf("%d + %d = %d\n", 3, 5, math.Add(3, 5))
	// fmt.Printf("%d - %d = %d\n", 5, 3, math.Sub(5, 3))
	// fmt.Printf("%d * %d = %d\n", 5, 3, math.Multiply(5, 3))
	// fmt.Printf("%d / %d = %.2f\n", 5, 3, math.Divide(5, 3))
	// fmt.Printf("%d / %d = %.2f\n", 5, 3, math.Divide(5, 3))
	// fmt.Printf("is %d even? - %v\n", 5, math.Even(5))
	// fmt.Printf("is %d odd? - %v\n", 5, math.Odd(5))
	// fmt.Printf("square root of %d is %.2f\n", 1, math.Root(1))
	// fmt.Printf("square root of %d is %.2f\n", 16, math.Root(16))
	// fmt.Printf("%dth power of %d is %d\n", 2, 2, math.Pow(2, 2))
	// fmt.Printf("%dth power of %d is %d\n", 100, 1, math.Pow(1, 100))
	// fmt.Printf("%dth power of %d is %d\n", 4, 2, math.Pow(2, 4))

	// Math(v2)
	fmt.Printf("%d + %d = %d\n", 3, 5, math.Add(3, 5))         // int
	fmt.Printf("%v + %v = %v\n", 3.5, 5.1, math.Add(3.5, 5.1)) // float
	fmt.Printf("%d - %d = %d\n", 5, 3, math.Sub(5, 3))         // int
	fmt.Printf("%v - %v = %v\n", 5.5, .2, math.Sub(5.5, .2))   // float
	fmt.Printf("%d * %d = %d\n", 5, 3, math.Multiply(5, 3))    // int
	fmt.Printf("%d * %d = %d\n", 5, 3, math.Multiply(5, 3))    // float
	fmt.Printf("%d / %d = %v\n", 5, 3, math.Divide(5, 3))      // int
	fmt.Printf("%v / %v = %v\n", .5, .3, math.Divide(.5, .3))  // float
	fmt.Printf("is %d even? - %v\n", 4, math.Even(4))          // int
	fmt.Printf("is %d odd? - %v\n", 3, math.Odd(3))            // int

	fmt.Printf("%dth power of %d is %d\n", 2, 2, math.Pow(2, 2))         // int
	fmt.Printf("%vth power of %v is %v\n", 2.5, 2.5, math.Pow(2.5, 2.5)) // float

	fmt.Printf("square root of %d is %v\n", 1, math.Root(1))       // int
	fmt.Printf("square root of %v is %v\n", 16.0, math.Root(16.0)) // float

	fmt.Printf("square root of %d is %v\n", 16, math.Root(16))     // int
	fmt.Printf("square root of %d is %v\n", 24, math.Root(24))     // int
	fmt.Printf("square root of %v is %v\n", 25.0, math.Root(25.0)) // float
}
