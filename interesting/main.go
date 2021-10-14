// package main

// import (
// 	"fmt"
// )

// func merge(out chan<- string, inA, inB <-chan string) {
// 	for {
// 		select {
// 		case v, ok := <-inA:
// 			if !ok{
// 				inA = nil
// 				if inB == nil {
// 					break
// 				}
// 			}
// 			out <- v
// 		case v := <-inB:
// 			out <- v
// 		}
// 	}
// }

// func main() {
// 	out := make(chan string)
// 	inA := make(chan string)
// 	inB := make(chan string)

// 	go merge(out, inA, inB)
// 	go func() {
// 		for i := 1; i < 10; i++ {
// 			inA <- "ss"
// 		}
// 		close(inA)
// 	}()
// 	go func() {
// 		for i := 1; i < 10; i++ {
// 			inB <- "ss"
// 		}
// 		close(inB)
// 	}()
// 	for v := range out {
// 		fmt.Print(v)

// 	}
// 	fmt.Println("All done")
// }
package main

import (
	"fmt"
	"sync"
)

func main() {
	i := 0
	N := 1000
	wg := sync.WaitGroup{}
	// mu := sync.Mutex{}
	wg.Add(N)
	for j := 1; j <= N; j++ {
		go func(i *int) {
			defer wg.Done()
			// mu.Lock()
			*i++
			// mu.Unlock()
		}(&i)
	}
	wg.Wait()
	fmt.Println(i)
}
