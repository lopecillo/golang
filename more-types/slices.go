package main

import (
	"fmt"
	"strings"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printStringAndSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func printTicTacToeBoard(board [][]string) {
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	fmt.Printf("\n")
}

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Slices
	var slice []int = primes[1:4]
	fmt.Println(slice)

	// Slices behave like pointers
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	slice_a := names[0:2]
	slice_b := names[1:3]
	fmt.Println(slice_a, slice_b)

	slice_b[0] = "XXX"
	fmt.Println(slice_a, slice_b)
	fmt.Println(names)

	// Slice literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	slice_literal := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(slice_literal)

	// Slice bounds
	slice_bounds := []int{2, 3, 5, 7, 11, 13}

	slice_bounds = slice_bounds[1:4]
	fmt.Println(slice_bounds)

	slice_bounds = slice_bounds[:2]
	fmt.Println(slice_bounds)

	slice_bounds = slice_bounds[1:]
	fmt.Println(slice_bounds)

	// Slice length and capacity

	// Slice the slice to give it zero length.
	slice_bounds = slice_bounds[:0]
	printSlice(slice_bounds)

	// Extend its length (only within its capacity limits)
	slice_bounds = slice_bounds[:4]
	printSlice(slice_bounds)

	// Drop its first two values.
	slice_bounds = slice_bounds[2:]
	printSlice(slice_bounds)

	// Nil slices
	var nil_slice []int
	fmt.Println(nil_slice, len(nil_slice), cap(nil_slice))
	if nil_slice == nil {
		fmt.Println("nil!")
	}

	// Creating slices with make
	a := make([]int, 5)
	printStringAndSlice("a", a)

	b := make([]int, 0, 5)
	printStringAndSlice("b", b)

	c := b[:2]
	printStringAndSlice("c", c)

	d := c[2:5]
	printStringAndSlice("d", d)

	// Slices of slices
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	printTicTacToeBoard(board)
	board[2][2] = "O"
	printTicTacToeBoard(board)
	board[1][2] = "X"
	printTicTacToeBoard(board)
	board[1][0] = "O"
	printTicTacToeBoard(board)
	board[0][2] = "X"
	printTicTacToeBoard(board)
	board[0][1] = "O"
	printTicTacToeBoard(board)
	board[1][1] = "X"
	printTicTacToeBoard(board)
	board[2][0] = "O"
	printTicTacToeBoard(board)
	board[2][1] = "X"
	printTicTacToeBoard(board)

	var slice_append []int
	printSlice(slice_append)

	// append works on nil slices.
	slice_append = append(slice_append, 0)
	printSlice(slice_append)

	// The slice grows as needed.
	slice_append = append(slice_append, 1)
	printSlice(slice_append)

	// We can add more than one element at a time.
	slice_append = append(slice_append, 2, 3, 4)
	printSlice(slice_append)
}
