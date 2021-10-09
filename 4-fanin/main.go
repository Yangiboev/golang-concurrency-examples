package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
func fanin(ch1, ch2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			v1 := <-ch1
			c <- v1
		}
	}()
	go func() {
		for {
			c <- <-ch1
		}
	}()
	return c
}
func faninSample(cs ...<-chan string) <-chan string {
	c := make(chan string)
	for _, v1 := range cs {
		go func(cv <-chan string) { // cv is a channel value
			for {
				c <- <-cv
			}
		}(v1) // send each channel to
	}
	return c
}
func main() {
	c := faninSample(boring("boring1!"), boring("boring2!"))
	for i := 0; i < 5; i++ {
		fmt.Println(<-c) // now we can read from 1 channel
	}
	fmt.Println("You're both boring. I'm leaving")

}
