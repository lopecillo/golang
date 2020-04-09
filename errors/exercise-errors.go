package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	var (
		z          float64 = x / 2
		accuracy   float64 = 1e-12
		iterations uint
	)
	for iterations = 0; math.Abs(z*z-x) > accuracy; iterations++ {
		z -= (z*z - x) / (2 * z)
	}
	fmt.Printf("Value after %v iterations with %v accuracy:\n", iterations, accuracy)
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
