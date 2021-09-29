package main

import (
	"fmt"
	"sync"
)

// Please describe what can be improved in the program?

const N = 10

func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i *int, m map[int]int) {
			defer wg.Done()
			m[*i] = *i + 1
		}(&i, m)
	}

	wg.Wait()

	fmt.Println(m)
}
