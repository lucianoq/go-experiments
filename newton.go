package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64 = 20.0
	for i := 0; i < 8; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func main() {
	var newton float64 = Sqrt(2)
	var inmath float64 = math.Sqrt(2)
	var diff float64 = newton - inmath
	fmt.Println("NEWTON: ", newton)
	fmt.Println("MATH: ", inmath)
	fmt.Println("DIFF: ", diff)
}
