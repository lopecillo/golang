package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	accuracy := 1e-12
	iterations := 0
	for iterations = 0; math.Abs(z*z - x) > accuracy; iterations++ {
		z -= (z*z - x) / (2*z)
	}
	fmt.Printf("Value after %v iterations with %v accuracy:\n", iterations, accuracy)
	return z
}

func main() {
	const num = 2
	fmt.Println("Input:", num)
	fmt.Println("Homemade:", Sqrt(num))
	fmt.Println("Math lib:", math.Sqrt(num))
}