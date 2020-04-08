package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground! The time is:")

	fmt.Println("[Raw]\t\t", time.Now())
	fmt.Println("[RFC822]\t", time.Now().Format(time.RFC822))
	fmt.Println("[RFC3339]\t", time.Now().Format(time.RFC3339))

	// Layout:
	//	Jan 2 15:04:05 2006 MST
	// An easy way to remember this value is that it holds, when presented
	// in this order, the values (lined up with the elements above):
	//	  1 2  3  4  5    6  -7
	// For more detail visit:
	// https://yourbasic.org/golang/format-parse-string-time-date-example/
	fmt.Println("[Custom layout]\t", time.Now().Format("2006-01-02 15:04:05 MST"))
}