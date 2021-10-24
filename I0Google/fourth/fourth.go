package main

func fanIn(msg1, msg2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-msg1:
				c <- s
			case s := <-msg2:
				c <- s
			}
		}
	}()
	return c 
}
