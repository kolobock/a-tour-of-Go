package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		fmt.Printf("iteration %v, sqrt %v\n", i, z)
		z_before := z
		z -= (z*z - x) / (2 * z)
		if math.Abs(z_before-z) < 0.0000000001 {
			return z
		}
	}
	return z
}

func main() {
	num := float64(2)
	fmt.Printf("mine Sqrt: %v\n", Sqrt(num))
	fmt.Printf("math Sqrt: %v\n", math.Sqrt(num))
}
