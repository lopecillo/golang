package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("My favourite number is", rand.Intn(10))
}