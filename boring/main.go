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
	<-time.After(time.Second * 2)
	fmt.Println("You are boring I am leaving!")

}
