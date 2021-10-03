package main

import "fmt"

func main() {
	ch, cancel := CountTo(100)
	for i := range ch {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
	cancel()
}
func CountTo(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})
	cancel := func() {
		close(done)
	}
	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				return
			case ch <- i:
			}
		}
		close(ch)
	}()
	return ch, cancel
}
