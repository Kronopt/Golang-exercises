// 1. Implement the Walk function.
// 2. Test the Walk function.
// The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.
// Create a new channel ch and kick off the walker: go Walk(tree.New(1), ch)
// Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.
// 3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.
// 4. Test the Same function.
// Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.
package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// ChooseBranch walks every branch
func ChooseBranch(t *tree.Tree, ch chan int) {
	if t != nil {
		ChooseBranch(t.Left, ch)
		ch <- t.Value
		ChooseBranch(t.Right, ch)
	}
	return
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ChooseBranch(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	t1Ch := make(chan int)
	t2Ch := make(chan int)
	go Walk(t1, t1Ch)
	go Walk(t2, t2Ch)
	for t1Value := range t1Ch {
		t2Value, ok := <-t2Ch
		if !ok || t1Value != t2Value {
			return false
		}
	}
	// t1Ch is closed for sure because the loop is over
	// check t2Ch
	_, ok := <-t2Ch
	if ok {
		return false
	} // not closed
	return true
}

func main() {
	ch := make(chan int)
	tr := tree.New(1)
	go Walk(tr, ch)
	for value := range ch {
		fmt.Print(value, ",")
	}
	fmt.Print("\n")
	fmt.Println(Same(tree.New(1), tree.New(1)), "= true")
	fmt.Println(Same(tree.New(1), tree.New(2)), "= false")
}
