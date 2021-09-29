package main

import "sync"

func processAndGather(in <-chan int, processor func(int) int, num int) []int {
	out := make(chan int, num)
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				out <- processor(v)
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	var result []int
	for v := range out {
		result = append(result, v)
	}
	return result
}
func main() {
	inputChannel := make(chan int)
	for i := 0; i < 100; i++ {
		inputChannel <- i
	}
	processAndGather(inputChannel, process, 10)

}

func process(i int) int {
	return i
}
