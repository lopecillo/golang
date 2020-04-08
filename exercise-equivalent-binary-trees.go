package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

//type Tree struct {
//    Left  *Tree
//    Value int
//    Right *Tree
//}

func WalkImpl(t *tree.Tree, ch, quit chan int) {
	if t == nil {
		return
	}
	WalkImpl(t.Left, ch, quit)
	select {
	case ch <- t.Value:
		// Value successfully sent.
	case <-quit:
		return
	}
	WalkImpl(t.Right, ch, quit)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch, quit chan int) {
	WalkImpl(t, ch, quit)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
// NOTE: An implementation without the "quit" channel
//       would leak goroutines when trees are different.
func Same(t1, t2 *tree.Tree) bool {
	w1, w2 := make(chan int), make(chan int)
	quit := make(chan int)

	go Walk(t1, w1, quit)
	go Walk(t2, w2, quit)

	for {
		v1, ok1 := <-w1
		v2, ok2 := <-w2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	fmt.Printf("tree.New(1) = %v\n", tree.New(1))
	fmt.Printf("tree.New(2) = %v\n", tree.New(2))

	fmt.Print("tree.New(1) == tree.New(1): ")
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}

	fmt.Print("tree.New(1) != tree.New(2): ")
	if !Same(tree.New(1), tree.New(2)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}
}