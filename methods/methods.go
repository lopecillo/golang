package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// This is a method for the type Vertex
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Methods with pointer receivers can modify the receiver's values
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Printf("Before scaling: %v, Abs: %v\n", v, v.Abs())

	v.Scale(10)
	fmt.Printf("After scaling: %v, Abs: %v\n", v, v.Abs())

	p := &Vertex{4, 3}
	fmt.Printf("With a pointer: %v, Abs: %v\n", *p, p.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}