package main

import (
	"fmt"
	"sync"
)

func main() {
	// 일반적으로 내컴퓨터에서는 고루틴이 실행되기도 전에 루프가 종료되므로,
	// salutation 은 문자열 슬라이스의 마지막 값인 "good day"에 대한 참조를 저장하고 있는
	// 힙으로 옮겨지게 되고, 이에 따라 보통은 "good day"가 세번 출력된다.
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greeings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
	// Output
	// good day x 3
}
