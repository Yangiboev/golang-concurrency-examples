package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var slice []int
// 	var wg = &sync.WaitGroup{}
// 	queue := make(chan int, 1)

// 	wg.Add(100)
// 	for i := 0; i < 100; i++ {
// 		go func(i int) {
// 			queue <- i
// 		}(i)
// 	}
// 	go func() {
// 		for v := range queue {
// 			slice = append(slice, v)
// 			wg.Done()
// 		}
// 	}()
// 	wg.Wait()
// 	fmt.Println(slice)
// }
// func main() {
// 	runtime.GOMAXPROCS(1)

// 	done := false

// 	go func() {
// 		done = true
// 	}()

// 	for !done {
// 	}
// 	fmt.Println("finished")
// }
