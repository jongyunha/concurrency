package main

import "fmt"

// plusOne return a channel of num + 1 for nums received from in
func plusOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// 생성자 패턴과 마찬가지로 데이터를 보내는 쪽에서 닫아줘야 합니다.
		defer close(out)
		for num := range in {
			out <- num + 1
		}
	}()
	return out
}

type intPipe func(<-chan int) <-chan int

func chain(ps ...intPipe) intPipe {
	return func(in <-chan int) <-chan int {
		c := in
		for _, p := range ps {
			c = p(c)
		}
		return c
	}
}

func main() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()
	// for num := range plusOne(plusOne(c)) {
	// fmt.Println(num)
	// Output:
	//   7
	//   5
	//   10
	// }

	plusTwo := chain(plusOne, plusOne, plusOne)
	for num := range plusTwo(c) {
		fmt.Println(num)
	}
}
