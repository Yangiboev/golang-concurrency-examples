package main

import (
	"fmt"
	"time"
)

func boring(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s - %d", msg, i):
			case <-quit:
				fmt.Println("clean up")
				quit <- "see you again"
				close(c)
				return
			}
			time.Sleep(time.Duration(1000) * time.Millisecond)
		}
	}()
	return c
}
func main() {
	quit := make(chan string)
	c := boring("Joe", quit)
	for i := 3; i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Bye"
	fmt.Println("Dell says:", <-quit)
}
