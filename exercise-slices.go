package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {

	array := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		array[y] = make([]uint8, dx)
	}

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			//array[y][x] = uint8((x+y)/2)
			//array[y][x] = uint8(x*y)
			array[y][x] = uint8(x^y)
		}
	}
	return array

}

func main() {
	pic.Show(Pic)
}