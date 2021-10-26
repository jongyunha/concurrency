package main

import "fmt"

func main() {
	go func() {
		// 메인 함수가 먼저 끝나서 수행되지 않음
		fmt.Println("In goroutine")
	}()
	fmt.Println("In main routine")
}
