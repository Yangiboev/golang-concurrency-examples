package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	msg       string
	waitForIT chan bool
}

func fanIn(inputs ...<-chan Message) <-chan Message {
	c := make(chan Message)
	for _, val := range inputs {
		input := val
		go func() {
			for {
				c <- <-input
			}
		}()
	}
	return c
}

func boring(msg string) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			c <- Message{
				msg:       fmt.Sprintf("%s - %d", msg, i),
				waitForIT: waitForIt,
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

			<-waitForIt
		}
	}()
	return c
}
func main() {
	c := fanIn(boring("Dell"), boring("ICon"))
	for i := 0; i < 5; i++ {
		msg := <-c
		fmt.Println(msg.msg)
		msg1 := <-c
		fmt.Println(msg1.msg)

		msg.waitForIT <- true
		msg1.waitForIT <- true

	}
	fmt.Println("You're both boring. I'm leaving")

}
