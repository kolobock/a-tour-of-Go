package main

import (
	"fmt"
	"sort"
	"time"

	"golang.org/x/tour/tree"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		go Walk(t.Left, ch)
	}

	if t.Right != nil {
		go Walk(t.Right, ch)
	}

	ch <- t.Value
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	fmt.Printf("t1: %v\nt2: %v\n", t1, t2)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	walk1 := readFromChan(ch1)
	walk2 := readFromChan(ch2)

	for i, v := range walk1 {
		if walk2[i] != v {
			return false
		}
	}

	return true
}

func readFromChan(ch chan int) (v []int) {
	for {
		select {
		case x := <-ch:
			v = append(v, x)
		default:
			time.Sleep(50 * time.Millisecond)
		}
		if len(v) == 10 {
			break
		}
	}
	sort.Ints(v)

	return
}

func main() {
	ch := make(chan int)
	t1 := tree.New(1)
	t2 := tree.New(2)
	fmt.Println("tree 1:", t1)
	go Walk(t1, ch)
	fmt.Printf("%d", readFromChan(ch))
	fmt.Println("")
	fmt.Println("")

	fmt.Printf("\tsame: %v\n", Same(t1, t1))                   // true
	fmt.Printf("\tsame: %v\n", Same(tree.New(1), tree.New(1))) // true
	fmt.Printf("\tsame: %v\n", Same(t1, t2))                   // false
	fmt.Printf("\tsame: %v\n", Same(tree.New(2), tree.New(2))) // true
	fmt.Printf("\tsame: %v\n", Same(tree.New(3), tree.New(3))) // true
}
