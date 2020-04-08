package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	var (
		a int = 42
		b = 13
	)
	fmt.Println("- \"add\" function:")
	fmt.Println(a, "+", b, "=", add(a, b))

	var (
		string1 string = "hello"
		string2 = "world"
	)
	fmt.Println("- \"swap\" function:")
	fmt.Println(string1, string2)
	fmt.Println("Swapped:")
	swap1, swap2 := swap(string1, string2)
	fmt.Println(swap1, swap2)

	fmt.Println("- \"split\" function (named return values with naked return):")
	fmt.Println(split(17))

}