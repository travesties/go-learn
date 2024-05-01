package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}
	walker(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	cht1 := make(chan int)
	cht2 := make(chan int)

	go Walk(t1, cht1)
	go Walk(t2, cht2)

	for {
		v1, ok1 := <-cht1
		v2, ok2 := <-cht2

		if v1 != v2 || ok1 != ok2 {
			return false
		}

		if !ok1 {
			break
		}
	}

	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
