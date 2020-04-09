package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	// Accessing fields
	v01 := Vertex{1, 2}
	v01.X = 4
	fmt.Println(v01.X)

	// Pointers to structs
	v02 := Vertex{1, 2}
	p02 := &v02
	p02.X = 1e9
	fmt.Println(v02)

	// Struct literals
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p  = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, p, v2, v3)
}