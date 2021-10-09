package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
func main() {
	go boring("I am boring!")
	fmt.Println("I am listenning!")
	time.Sleep(2 * time.Second)
	fmt.Println("You are boring I am leaving!")

}
