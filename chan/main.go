package main

import (
	"fmt"
	"time"
)

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(1e3) * time.Millisecond)
	}
}
func main() {
	c := make(chan string)
	go boring("boring!", c)

	for i := 0; i < 5; i++ {
		fmt.Printf("You say %q\n", <-c)
	}
	fmt.Println("DONW!")
}
