package main

import (
	"fmt"
	"sync"
)

// Clicked 라는 조건을 가지고 있는 Button 타입을 정의한다.
type Button struct {
	Clicked *sync.Cond
}

func main() {
	button := Button{
		Clicked: sync.NewCond(&sync.Mutex{}),
	}

	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)

		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	// Mouse 버튼이 클릭되고 난뒤의 핸들러를 설정한다. 이는 결국 Clicked cond 에서
	// Broadcast 를 호출해 모든 핸들러에게 마우스 버튼이 클릭 됐음을 알린다.
	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	// Broadcast 를 한번 호출하면 모든 핸들러가 모두 실행된다.
	button.Clicked.Broadcast()
	clickRegistered.Wait()
}
