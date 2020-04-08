package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	// Calling a method on a nil interface is a run-time error
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}