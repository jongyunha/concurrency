package main

import "fmt"

func main() {
	// bad pattern
	// c := make(chan int)
	// go func() {
	//   c <- 1
	//   c <- 2
	//   c <- 3
	//   close(c)
	// }()
	//
	// for num := range c {
	//   fmt.Println(num)
	// }

	// good pattern
	// 단방향 채널을 반환하여 이 채널을 이용하는 고루틴이 받아 가기만 할 수 있게 제한
	// 그렇지 않을경우 이채널에 값을 보내려고 시도하면 그자료를 받아가는 고루틴이 없어서 영원히 프로그램이 멈출수 있다
	c2 := func() <-chan int {
		c2 := make(chan int)
		go func() {
			defer close(c2)
			c2 <- 1
			c2 <- 2
			c2 <- 3
		}()
		return c2
	}()

	for num := range c2 {
		fmt.Println(num)
	}
}
