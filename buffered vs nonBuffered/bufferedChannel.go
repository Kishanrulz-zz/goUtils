package main

import (
	"fmt"
	"time"
)

// Buffered Channel is asynchronous
// We can write to a buffered channel till the buffer is full
// This makes buffered channel non blocking
func main() {
	message := make(chan string,1) //  Buffered Channel
	count := 6

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