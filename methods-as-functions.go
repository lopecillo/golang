package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Printf("Before scaling: %v, Abs: %v\n", v, Abs(v))

	Scale(&v, 10)
	fmt.Printf("After scaling: %v, Abs: %v\n", v, Abs(v))

	p := &Vertex{4, 3}
	fmt.Printf("With a pointer: %v, Abs: %v\n", *p, Abs(*p))
}