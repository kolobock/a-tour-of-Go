package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		fmt.Printf("iteration %v, sqrt %v\n", i, z)
		z_before := z
		z -= (z*z - x) / (2 * z)
		if math.Abs(z_before-z) < 0.0000000001 {
			return z, nil
		}
	}
	return z, nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func main() {
	nums := []float64{2, -2}
	for _, num := range nums {
		fmt.Printf("** Sqrt for %v\n", num)
		res, err := Sqrt(num)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("mine Sqrt: %v\n", res)
		fmt.Printf("math Sqrt: %v\n\n", math.Sqrt(num))
	}
}
