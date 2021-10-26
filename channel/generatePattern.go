package main

import "fmt"

func fibonacci(max int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		a, b := 0, 1
		for a <= max {
			c <- a
			a, b = b, a+b
		}
	}()
	fmt.Printf("%T", c)
	return c
}

func fibonacciClosure(max int) func() int {
	next, a, b := 0, 0, 1
	return func() int {
		next, a, b = a, b, a+b
		if next > max {
			return -1
		}
		return next
	}
}

func main() {
	for fib := range fibonacci(15) {
		fmt.Printf("%d, ", fib)
	}

	fmt.Println()

	fib := fibonacciClosure(15)
	for n := fib(); n >= 0; n = fib() {
		fmt.Printf("%d, ", n)
	}
}
