package main

import (
	"fmt"
	"time"
)


// Non buffered channels are synchronous.
// Until and unless an object has been read from the channel, we cannot put more values into the channel.
// This makes non-buffered channel blocking
func main() {
	message := make(chan string) // non buffered channel
	count := 3

	go func() {
		for i := 1; i <= count; i++ {
			fmt.Println("send message")
			message <- fmt.Sprintf("message %d", i)
		}
	}()

	time.Sleep(time.Second * 3)

	for i := 1; i <= count; i++ {
		fmt.Println(<-message)
	}
}

