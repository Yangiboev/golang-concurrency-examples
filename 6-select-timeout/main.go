package main

import (
	"fmt"
	"time"
)

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s - %d", msg, i)
			time.Sleep(time.Duration(1500) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	c := boring("Dell")
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("We talked too much.")
			return
		}
	}
}
