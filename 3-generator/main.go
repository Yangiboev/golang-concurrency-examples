package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 1; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}
func main() {
	joe := boring("dell")
	icon := boring("icon")
	for i := 0; i < 10; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-icon)
	}
	fmt.Println('A')
	fmt.Println('a' - 1)
	fmt.Println("You are both boring. I am leaving")
}
