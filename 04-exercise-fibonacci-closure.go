package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fibo_2, fibo_1 := 0, 1

	return func() int {
		f := fibo_2
		fibo_2, fibo_1 = fibo_1, fibo_1+f

		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
