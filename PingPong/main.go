package main

import "fmt"

func ping(c chan<- string, msg string) {
	c <- msg
}
func pong(ping <-chan string, pong chan<- string) {
	msg := <-ping
	pong <- msg
}
func main() {
	pingc := make(chan string, 1)
	pongc := make(chan string, 1)
	ping(pingc, "helllo")
	pong(pingc, pongc)
	fmt.Println(<-pongc)
}
