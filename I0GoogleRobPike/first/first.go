package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)
	go boring("boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You would say %q \n", <-c)
	}
}

func boring(str string, ch chan string) {
	for i := 0; ; i++ {
		ch <- fmt.Sprintf("%s %d", str, i)
		time.Sleep(time.Duration(rand.Int63n(1e3)) * time.Millisecond)
	}
}
