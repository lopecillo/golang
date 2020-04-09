package main

import "fmt"

var power_two = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range power_two {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow := make([]int, 10)
	// Skip index or value
	// for i, _ := range pow
	// for _, value := range pow
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}