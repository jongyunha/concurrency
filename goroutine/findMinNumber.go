package main

import (
	"fmt"
	"sync"
)

func min(a []int) int {
	if len(a) == 0 {
		return 0
	}

	min := a[0]
	for _, e := range a[1:] {
		if min > e {
			min = e
		}
	}
	return min
}

// 병렬 버전으로 가장 작은값 찾기
// 사람 4명이서 가장 작은수를 찾는다고 생각해보면, 전체를 넷으로
// 나눠서 각자 가장 작은수를 찾은다음에 그중에서 가장 작은수를 한번더 찾는다.
func parallelmin(a []int, n int) int {
	if len(a) < n {
		return min(a)
	}

	mins := make([]int, n)
	size := (len(a) + n - 1) / n
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			begin, end := i*size, (i+1)*size
			if end > len(a) {
				end = len(a)
			}
			mins[i] = min(a[begin:end])
		}(i)
	}
	wg.Wait()
	return min(mins)
}

func main() {
	nums := []int{
		83, 46, 49, 17, 92,
		68, 39, 91, 44, 99,
		25, 42, 74, 56, 23,
	}
	fmt.Println(parallelmin(nums, 4))
}
